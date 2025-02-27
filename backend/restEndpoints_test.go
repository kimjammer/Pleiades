package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func connectDB() {
	//Connect to database
	if db != nil {
		return
	}

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

		//Can't defer db disconnection because it will close the connection before the tests are run
		//log.Println("Successfully connected to MongoDB")
		//defer func() {
		//	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		//		panic(err)
		//	}
		//}()
	}
}

func resetDB() {
	db.Drop(context.TODO())

	//Create a user
	user := User{
		Id:          "c8478bcd-51fe-486d-965d-c8a8837c577c",
		FirstName:   "Joe",
		LastName:    "Smith",
		Email:       "example@example.com",
		Password:    "notthepassword",
		PhoneNumber: "1234567890",
		Projects:    []string{""},
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

func TestNewProjectNoBody(t *testing.T) {
	router := setupRouter()
	connectDB()
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewProjectEmptyTitle(t *testing.T) {
	router := setupRouter()
	connectDB()
	resetDB()

	request := NewProjectRequest{
		Title:       "",
		Description: "",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewProjectSuccess(t *testing.T) {
	router := setupRouter()
	connectDB()
	resetDB()

	request := NewProjectRequest{
		Title:       "Test Project",
		Description: "Test Description",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
