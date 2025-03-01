package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//Connect to database
	godotenv.Load()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalln("Set your 'MONGODB_URI' environment variable as in .env.example. Aborting tests.")
	} else {
		mongoClient, err := mongo.Connect(options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		db = mongoClient.Database("pleiades-test")

		log.Println("Successfully connected to MongoDB")
		defer func() {
			if err := mongoClient.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	}

	//Run tests
	m.Run()
}

// Util to reset the DB to a known state for testing
func resetDB() {
	_ = db.Drop(context.TODO())

	//Create a user
	user := User{
		Id:          "c8478bcd-51fe-486d-965d-c8a8837c577c",
		FirstName:   "Joe",
		LastName:    "Smith",
		Email:       "example@example.com",
		Password:    "notthepassword",
		PhoneNumber: "1234567890",
		Projects:    []string{"53ed4d28-9279-4b4e-9256-b1e693332625"},
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	project := Project{
		Id:          "53ed4d28-9279-4b4e-9256-b1e693332625",
		Title:       "Test Project",
		Description: "Test Description",
		Users:       []string{"c8478bcd-51fe-486d-965d-c8a8837c577c"},
		Tasks:       []Task{},
		Polls:       []Poll{},
	}
	_, _ = db.Collection("projects").InsertOne(context.TODO(), project)
	_, _ = db.Collection("users").InsertOne(context.TODO(), user)
}

// Manually set the token for testing
func tokenOverride() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userId", "c8478bcd-51fe-486d-965d-c8a8837c577c")
		c.Next()
	}
}

// Setup router for testing with the token manually set
func setupTestRouter() *gin.Engine {
	// Setup webserver
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:4173"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(tokenOverride())
	defineRoutes(router)

	return router
}
