package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

/**
 * Iterate over every document in a collection and sum the length of the specified field.
 * Sets Gin status to server error if error encountered
 */
func countArrayField(ctx *gin.Context, collection, field string) int64 {
	cursor, err := db.Collection(collection).Aggregate(ctx, mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "total", Value: bson.D{{
				Key: "$sum",
				Value: bson.D{{
					Key: "$size",
					Value: "$" + field,
				}},
			}}},
		}}},
	})
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return 0
	}
	defer cursor.Close(ctx)

	var r struct{ Total int64 }
	if cursor.Next(ctx) && cursor.Decode(&r) == nil {
		ctx.Status(http.StatusInternalServerError)
		return r.Total
	}
	ctx.Status(http.StatusInternalServerError)
	return 0
}
