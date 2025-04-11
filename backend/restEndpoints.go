package main

import (
	"encoding/base64"
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
		Created:     int(time.Now().UnixMilli()),
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
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "projects", Value: crrUser.Projects}}}}
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
	newUser.Projects = []string{}
	//Insert new user into db
	collection := db.Collection("users")
	_, err = collection.InsertOne(c.Request.Context(), newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	log.Println("Created user with ID: ", newUser.Id)
	//set cookie upon successful registration (cookie is user id)
	token := makeToken(newUser.Id.Hex())
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600, "/", "", true, true)
	log.Println("Created User with Id: ", newUser.Id)
	c.JSON(http.StatusCreated, gin.H{"success": true, "userId": newUser.Id})
}

func login(c *gin.Context) {
	listUsers()

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
			log.Println("Token:", token)
			c.SetCookie("token", token, 3600, "/", "", true, true)
			c.JSON(http.StatusOK, gin.H{"exists": true, "userId": user.Id})
		} else {
			c.JSON(http.StatusOK, gin.H{"exists": false})
		}
	}

}

func invite(c *gin.Context) {
	//Get current user
	crrUser, err := getUser(c)
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Validate permissions (is project member)
	projectId := c.Query("id")
	isMember := slices.Contains(crrUser.Projects, projectId)
	if !isMember {
		// TODO: convert to middleware
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a project member"})
		return
	}

	invitations := db.Collection("invitations")

	// Create 1 week ttl TODO: is this ok to do every time or is there a better init location?
	ttl := int32(604800) // 1 week, in seconds
	if TEST {
		ttl = 20 // seconds
	}
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "createdat", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(ttl).SetName("ttl_index"),
	}
	invitations.Indexes().CreateOne(c, indexModel)

	// Create record
	invitation := Invitation{
		Id:        uuid.New().String(),
		CreatedAt: time.Now(),
		ProjectId: projectId,
	}
	_, err = invitations.InsertOne(c, invitation)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, invitation.Id)
}

func join(c *gin.Context) {
	// get invite entry
	invitation, err := getInvite(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invite not found"})
		return
	}

	userId := c.GetString("userId")

	// Add user to project
	err = updateProject(invitation.ProjectId, func(project *Project) error {
		for i, user := range project.Users {
			if user.User == userId {
				project.Users[i].LeftProject = false
				return nil
			}
		}

		project.Users = append(project.Users, UserAndLeft{User: userId, LeftProject: false})

		return nil
	})
	if err != nil {
		panic("The updater cannot return an error")
	}

	requeryUsersForProject(invitation.ProjectId)

	// Add project to user
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}, {Key: "projects", Value: bson.M{"$ne": invitation.ProjectId}}}
	update := bson.M{"$push": bson.M{"projects": invitation.ProjectId}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.String(http.StatusOK, "success")
}

func joinInfo(c *gin.Context) {
	invitation, err := getInvite(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"exists": false})
		return
	}

	// Get addl project info TODO: abstract
	filter := bson.D{{Key: "_id", Value: invitation.ProjectId}}
	var project Project
	err = db.Collection("projects").FindOne(c, filter).Decode(&project)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"exists": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": true, "id": invitation.ProjectId, "title": project.Title, "description": project.Description})
}

func setAvailability(c *gin.Context) {
	// Decode new availability
	var newAvailability []Availability
	if err := c.ShouldBindJSON(&newAvailability); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Write new availability to db
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.M{"$set": bson.M{"availability": newAvailability}}
	_, err := db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		panic(err)
	}

	user, _ := getUser(c) // error not possible b/c already validated

	requeryUser(user)
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
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.Status(http.StatusOK)
}

// Used by frontend to check if user has a valid session, since this is an authRequired() endpoint
func verifySession(c *gin.Context) {
	c.Status(http.StatusOK)
}

// used to get basic user info such as name
func getUserInfo(c *gin.Context) {
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

	log.Println("User name:", crrUser.FirstName)
	c.JSON(http.StatusOK, gin.H{"firstname": crrUser.FirstName})

}

func uploadProfilePic(c *gin.Context) {
	log.Println("uploading profile pic")
	//get user
	crrUser, err := getUser(c)
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	//log.Println("crrUser: ", crrUser)

	//store profile pic
	var reqBody struct {
		Image string `json:"image"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Decode the base64 string to binary data
	imgData, err := base64.StdEncoding.DecodeString(reqBody.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64"})
		return
	}
	crrUser.UserPhoto = imgData
	filter := bson.M{"_id": crrUser.Id}
	update := bson.M{"$set": bson.M{"userPhoto": imgData}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	log.Println("profile pic uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func getProfilePic(c *gin.Context) {
	//Get user
	crrUser, err := getUserById(c, c.Query("id"))
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err != nil { //if no pic
		log.Println("error encountered")
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile picture not found"})
		return
	} else if crrUser.UserPhoto == nil {
		log.Println("no profile pic in database")
		c.JSON(http.StatusOK, gin.H{"found": false})
	}
	log.Println("returning profile pic")
	c.Header("Content-Type", "image/png")
	c.Writer.Write(crrUser.UserPhoto)
}

func fetchName(c *gin.Context) {
	//Get user
	crrUser, err := getUserById(c, c.Query("id"))
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"firstName": crrUser.FirstName, "lastName": crrUser.LastName})
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
func deleteAllProjects() {
	collection := db.Collection("projects") // Reference the "users" collection

	result, err := collection.DeleteMany(context.TODO(), bson.M{}) // Empty filter {} means delete all
	if err != nil {
		log.Fatal("Error deleting projects:", err)
	}

	log.Printf("Deleted %d users from the database\n", result.DeletedCount)
}
