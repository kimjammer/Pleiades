package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func projectsHandler(c *gin.Context) {
	//Get current user
	//TODO: Get User id from token
	userId := "BOOTSTRAPPER"

	filter := bson.D{{Key: "_id", Value: userId}}
	var crrUser User
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//TODO: Add error handling for user not found
		} else {
			panic(err)
		}
	}

	//Get projects
	var response ProjectsResponse
	for _, projectId := range crrUser.Projects {
		filter := bson.D{{Key: "_id", Value: projectId}}
		var project Project
		err := db.Collection("projects").FindOne(c, filter).Decode(&project)
		if err != nil {
			panic(err)
		}
		response.Projects = append(response.Projects, minimalProject{
			Id:          project.Id,
			Title:       project.Title,
			Description: project.Description,
		})
	}

	c.JSON(http.StatusOK, response)
}

func newProjectHandler(c *gin.Context) {
	log.Println("new project")
	//Get current user
	//TODO: Get User id from token
	userId := "BOOTSTRAPPER"
	filter := bson.D{{"_id", userId}}

	var crrUser User
	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//TODO: Remove bootstrapper
			tempBootstrapper()
			err = db.Collection("users").FindOne(context.TODO(), filter).Decode(&crrUser)
			//TODO: Add error handling for user not found
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		}
	}

	var newProject NewProjectRequest
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	//Validation
	if newProject.Title == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	//Create Project
	project := Project{
		Id:          uuid.New().String(),
		Title:       newProject.Title,
		Description: newProject.Description,
		Users:       []string{crrUser.Id},
		Tasks:       []Task{},
		Polls:       []Poll{},
	}
	_, err = db.Collection("projects").InsertOne(c, project)
	if err != nil {
		log.Println("Error inserting project: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	//Add project to user
	crrUser.Projects = append(crrUser.Projects, project.Id)
	update := bson.D{{"$set", bson.D{{"projects", crrUser.Projects}}}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		log.Println("Error updating user with new project: ", err)
		c.Status(http.StatusInternalServerError)
		return
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	result, err := db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}

	log.Println("Inserted User with Id: ", result)
}

// this function checks to see if an email is already registered in the database
func checkEmail(c *gin.Context) {
	log.Println("checking email")
	email := c.Query("email")
	collection := db.Collection("users")
	var result bson.M

	//use c instead of context.TODO()???
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"exists": false})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"exists": true})
	}
	//c is sent at end of function
}

func registerUser(c *gin.Context) {
	var newUser User
	log.Println("registering new user")
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	} else {
		newUser.Password = string(hashedPassword)
	}
	//Insert the new user into MongoDB
	log.Println("inserting")
	collection := db.Collection("users")
	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func login(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	var result bson.M

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	log.Println(string(hashedPassword))
	log.Println("fetching from database")
	//use c instead of context.TODO()???
	err = db.Collection("users").FindOne(context.TODO(), bson.M{"email": email, "password": string(hashedPassword)}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"exists": false})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"exists": true})
	}
}
