package main

import (
	"context"
	"errors"
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
)

func getUserById(c context.Context, userId string) (crrUser User, err error) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	err = db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// no-op
		} else {
			panic(err)
		}
	}
	return
}

func getUser(c *gin.Context) (crrUser User, err error) {
	userId := c.GetString("userId")
	crrUser, err = getUserById(c, userId)
	return
}

func findUserByEmail(c *gin.Context, email string) (user User, err error) {
	log.Println("fetching from database")
	err = db.Collection("users").FindOne(c, bson.M{"email": email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"exists": false})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}
	return
}

/**
 * Get a project invitation object from database
 */
func getInvite(c *gin.Context) (invitation Invitation, err error) {
	joinId := c.Query("id")
	filter := bson.D{{Key: "_id", Value: joinId}}
	err = db.Collection("invitations").FindOne(c, filter).Decode(&invitation)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// No op
		} else {
			panic(err)
		}
	}
	return
}

func createInvite(c *gin.Context) string {
	//Get current user
	crrUser, err := getUser(c)
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return ""
	}

	// Validate permissions (is project member)
	projectId := c.Query("id")
	isMember := slices.Contains(crrUser.Projects, projectId)
	if !isMember {
		// TODO: convert to middleware
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a project member"})
		return ""
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

	return invitation.Id
}

func createUser(c *gin.Context, newUser *User) {
	log.Println("registering new user")
	newUser.NotifSettings = []bool{true, true, true}
	log.Println("Registering User, NotifSettings: ", newUser.NotifSettings)

	newUser.Id = primitive.NewObjectID()
	newUser.Availability = []Availability{}
	newUser.Projects = []string{}
	//Insert new user into db
	collection := db.Collection("users")
	_, err := collection.InsertOne(c, newUser)
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
