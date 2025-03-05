package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func projectsHandler(c *gin.Context) {
	//Get current user
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
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
	log.Println("Creating New Project")
	//Get current user
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}

	var crrUser User
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		//TODO: Add error handling for user not found
		log.Println("User not found!")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var newProject NewProjectRequest
	if err := c.ShouldBindJSON(&newProject); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	//Validation
	if newProject.Title == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Create Project
	project := Project{
		Id:          uuid.New().String(),
		Title:       newProject.Title,
		Description: newProject.Description,
		Users:       []UserAndLeft{{crrUser.Id.Hex(), false}},
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
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// this function checks to see if an email is already registered in the database
func checkEmail(c *gin.Context) {
	log.Println("checking email")
	email := c.Query("email")
	collection := db.Collection("users")
	var result bson.M

	err := collection.FindOne(c.Request.Context(), bson.M{"email": email}).Decode(&result)

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

	hashedPassword, err := encryptPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	} else {
		newUser.Password = hashedPassword
	}

	newUser.Id = primitive.NewObjectID()
	newUser.Availability = []Availability{}
	//Insert new user into db
	collection := db.Collection("users")
	_, err = collection.InsertOne(c.Request.Context(), newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
	//set cookie upon successful registration (cookie is user id)
	token := makeToken(newUser.Id.Hex())
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600, "/", "", true, true)
	log.Println("Created User with Id: ", newUser.Id)
}

func login(c *gin.Context) {
	//listUsers()
	email := c.Query("email")
	password := c.Query("password")
	var user User

	log.Println("fetching from database")
	err := db.Collection("users").FindOne(c.Request.Context(), bson.M{"email": email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"exists": false})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	} else {
		//check if passwords match
		if checkPassword(user.Password, password) {
			//set cookie upon successful login (cookie is user id)
			log.Println("User id: ", user.Id, user.Id.Hex())
			token := makeToken(user.Id.Hex())
			c.SetSameSite(http.SameSiteNoneMode)
			c.SetCookie("token", token, 3600, "/", "", true, true)
			c.JSON(http.StatusOK, gin.H{"exists": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"exists": false})
		}
	}

}

func userExists(userId string, c context.Context) (crrUser *User) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		} else {
			panic(err)
		}
	}
	return
}

func invite(c *gin.Context) {
	//Get current user
	projectId := c.Query("id")
	userId := c.GetString("userId")
	crrUser := userExists(userId, c);
	if crrUser == nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	log.Println(crrUser.FirstName)

	// Validate permissions (is project member)
	isMember := slices.Contains(crrUser.Projects, projectId)
	if !isMember {
		// TODO: convert to middleware
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a project member"})
		return
	}

	invitations := db.Collection("invitations")

	// Create 1 week ttl TODO: is this ok to do every time or is there a better init location?
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"CreatedAt", 1}},
		Options: options.Index().SetExpireAfterSeconds(604800).SetName("ttl_index"),
	}
	invitations.Indexes().CreateOne(c, indexModel)

	// Create record
	invitation := Invitation{
		Id:          uuid.New().String(),
		CreatedAt:   time.Now(),
	}
	_, err := invitations.InsertOne(c, invitation)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, invitation.Id)
}

func encryptPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Use a fixed cost
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// Reset token cookie
func logout(c *gin.Context) {
	//TODO: Also remove session from db
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.Status(http.StatusOK)
}

// Used by frontend to check if user has a valid session, since this is an authRequired() endpoint
func verifySession(c *gin.Context) {
	c.Status(http.StatusOK)
}

// TODO: Frontend e2e testing uses this button/route. Figure out solution for this before removing
func fakeLogin(c *gin.Context) {
	var user User
	_ = db.Collection("users").FindOne(c.Request.Context(), bson.D{}).Decode(&user)

	token := makeToken(user.Id.Hex())
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600, "/", "", true, true)
	c.Status(http.StatusOK)
}

// TEMPORARY
func listUsers() {
	collection := db.Collection("users") // Reference the "users" collection

	// MongoDB filter (empty filter `{}` to get all documents)
	filter := bson.M{}

	// Cursor to iterate over all documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error finding users:", err)
	}
	defer cursor.Close(context.TODO()) // Close cursor when function ends

	fmt.Println("Current Users in Database:")

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}
		fmt.Printf("ID: %s, Name: %s %s, Email: %s, Password: %s\n", user.Id, user.FirstName, user.LastName, user.Email, user.Password)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error:", err)
	}
}

// TEMPORARY
func deleteAllUsers() {
	collection := db.Collection("users") // Reference the "users" collection

	result, err := collection.DeleteMany(context.TODO(), bson.M{}) // Empty filter {} means delete all
	if err != nil {
		log.Fatal("Error deleting users:", err)
	}

	log.Printf("Deleted %d users from the database\n", result.DeletedCount)
}
