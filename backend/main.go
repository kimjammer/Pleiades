package main

import (
	"context"
	"log"
	"slices"

	// "net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database
var mailjetClient *mailjet.Client
var allowedOrigins = []string{"http://localhost:5173", "http://localhost:4173", "https://pleiades.pages.dev", "https://ethandawes.github.io"}

func shouldAllowOrigin(c *gin.Context, origin string) bool {
	// Must make special CORS exception for "log in with Google" because cross-origin redirects have `Origin: null`
	// From https://stackoverflow.com/questions/76085477/oauth2-after-redirect-response-request-origin-is-null#comment134184742_76085477
	// and my own testing
	// I could've just used js callback, but I thought this was more elegant b/c fewer steps
	var allowedOAuthRoutes = []string{"/login/google", "/register/google"}
	return slices.Contains(allowedOrigins, origin) || (origin == "null" && slices.Contains(allowedOAuthRoutes, c.Request.URL.Path))
}

func setupRouter() *gin.Engine {
	// Setup webserver
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOriginWithContextFunc = shouldAllowOrigin
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
	router.POST("/register/google", googleRegistration)
	router.GET("/login", login)
	router.POST("/login/google", googleLogin)
	router.POST("/logout", logout)
	router.GET("/verifySession", authRequired(), verifySession)
	router.GET("/invite", authRequired(), invite)
	router.GET("/join", authRequired(), join)
	router.GET("/join/info", joinInfo)
	router.POST("/availability", authRequired(), setAvailability)
	router.GET("/userInfo", authRequired(), getUserInfo) //TODO: remove this
	router.POST("/profilepic", authRequired(), uploadProfilePic)
	router.GET("/getprofilepic", authRequired(), getProfilePic)
	router.GET("/fetchName", authRequired(), fetchName)
	router.GET("/event", handleEvent)
	router.GET("/stats", reportStats)
	router.POST("/forgotPassword", forgotPasswordHandler)
	router.POST("/resetPassword", resetPasswordHandler)
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

	//Creat mailjet client
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")
	if (publicKey != "" && secretKey != "") || TEST {
		mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)
	} else {
		log.Println("No Mailjet API keys found, email sending disabled.")
	}

	router := setupRouter()
	defineRoutes(router)
	if TEST {
		defineTestRoutes(router)
	}
	router.Run(":8080")
}
