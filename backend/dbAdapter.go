package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

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