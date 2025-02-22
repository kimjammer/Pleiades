package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func projectsHandler(c *gin.Context) {
	//DEBUG: Dummy data
	response := ProjectsResponse{
		Projects: []minimalProject{
			{
				Id:          uuid.New().String(),
				Title:       "Project 1",
				Description: "This is project 1",
			},
		},
	}

	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func newProjectHandler(c *gin.Context) {
	//Get current user
	//TODO: Get User id from token
	filter := bson.D{{"_id", "BOOTSTRAPPER"}}

	var crrUser User
	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//TODO: Remove bootstrapper
			tempBootstrapper()
			err = db.Collection("users").FindOne(context.TODO(), filter).Decode(&crrUser)
			//TODO: Add error handling for user not found
		} else {
			panic(err)
		}
	}

	//Create Project
	project := Project{
		Id:          uuid.New().String(),
		Title:       "Project Title",
		Description: "Project Description",
		Users:       []string{crrUser.Id},
		Tasks:       []Task{},
		Polls:       []Poll{},
	}
	_, err = db.Collection("projects").InsertOne(context.TODO(), project)
	if err != nil {
		panic(err)
	}

	//Add project to user
	crrUser.Projects = append(crrUser.Projects, project.Id)
	update := bson.D{{"$set", bson.D{{"projects", crrUser.Projects}}}}
	_, err = db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

// TODO: Remove bootstrapping
func tempBootstrapper() {
	//Create a user
	user := User{
		Id:           "BOOTSTRAPPER",
		FirstName:    "Joe",
		LastName:     "Smith",
		Email:        "example@example.com",
		Password:     "notapassword",
		PhoneNumber:  "1234567890",
		UserPhoto:    "",
		Availability: nil,
		Projects:     nil,
	}

	result, err := db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}

	log.Println("Inserted User with Id: ", result)
}
