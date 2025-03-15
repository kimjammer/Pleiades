package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

// Manually set the token for testing
func tokenOverride() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userId", "67c7b20021675682e4395270")
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
