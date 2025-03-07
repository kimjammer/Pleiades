package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getUser(c *gin.Context) (crrUser User, err error) {
	userId := c.GetString("userId")
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