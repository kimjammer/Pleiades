package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"net/http"
	"os"
)

var db *mongo.Database

// Middleware to load token from cookie
func loadToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			log.Println("No token found")
		} else {
			userId, err := verifyToken(cookie)

			if err == nil {
				c.Set("userId", userId)
			}
		}

		c.Next()
	}
}

// Middleware to deny access to unauthenticated users, for routes that are only for logged-in users
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("userId")
		if exists == false {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	// Setup webserver
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:4173"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(loadToken())

	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/ws", authRequired(), wsEndpoint)
	router.GET("/projects", authRequired(), projectsHandler)
	router.POST("/projects/new", authRequired(), newProjectHandler)
	router.GET("/register/check", checkEmail)
	router.POST("/register/new", registerUser)
	router.GET("/login", login)

	//TODO: Remove testing route that only sets cookie
	router.GET("/fakelogin", fakeLogin)
}

func main() {
	//Connect to database
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("Set your 'MONGODB_URI' environment variable as in .env.example. Skipping database connection.")
	} else {
		mongoClient, err := mongo.Connect(options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		var result bson.M
		if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
			panic(err)
		}
		db = mongoClient.Database("pleiades")
		log.Println("Successfully connected to MongoDB")
		defer func() {
			if err := mongoClient.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	}

	router := setupRouter()
	defineRoutes(router)
	router.Run(":8080")
}
