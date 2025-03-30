package main

import (
	"context"
	"log"

	// "net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database
var allowedOrigins = []string{"http://localhost:5173", "http://localhost:4173", "https://pleiades.pages.dev", "https://ethandawes.github.io"}

func setupRouter() *gin.Engine {
	// Setup webserver
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(loadToken())

	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/ws", wsEndpoint)
	router.GET("/projects", authRequired(), projectsHandler)
	router.POST("/projects/new", authRequired(), newProjectHandler)
	router.GET("/register/check", checkEmail)
	router.POST("/register/new", registerUser)
	router.GET("/login", login)
	router.POST("/logout", logout)
	router.GET("/verifySession", authRequired(), verifySession)
	router.GET("/invite", authRequired(), invite)
	router.GET("/join", authRequired(), join)
	router.GET("/join/info", joinInfo)
	router.POST("/availability", authRequired(), setAvailability)
	router.GET("/userInfo", authRequired(), getUserInfo) //TODO: remove this
	router.POST("/profilepic", authRequired(), uploadProfilePic)
	router.GET("/getprofilepic", authRequired(), getProfilePic)
}

func defineTestRoutes(router *gin.Engine) {
	router.GET("/resetDatabase", resetDatabase)
}

func main() {
	log.Println("Test extensions enabled?", TEST)
	dbName := "pleiades"
	if TEST {
		dbName = "pleiades-test"
	}

	//Connect to database
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("MONGODB_URI environment variable not set!")
	} else {
		mongoClient, err := mongo.Connect(options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		var result bson.M
		if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
			panic(err)
		}
		db = mongoClient.Database(dbName)
		log.Println("Successfully connected to MongoDB")
		defer func() {
			if err := mongoClient.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	}

	router := setupRouter()
	defineRoutes(router)
	if TEST {
		defineTestRoutes(router)
	}
	router.Run(":8080")
}
