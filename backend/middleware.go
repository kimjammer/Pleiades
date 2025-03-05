package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Middleware to load token from cookie
func loadToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			log.Println("No token found")
		} else {
			userId, err := verifyToken(cookie)

			log.Println("Verfied token: ", userId)

			if err == nil {
				c.Set("userId", userId)
			}
		}

		c.Next()
	}
}

// Middleware to deny access to unauthenticated users, for routes that are only for logged-in users
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("userId")
		if exists == false {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}

/**
 * Middleware to set the 'invitation' context variable
 * Usage:
 * a, _ := c.Get("invitation")
 * b, _ := a.(Invitation)
 */
func getInvite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invitation Invitation
		joinId := c.Query("id")
		filter := bson.D{{Key: "_id", Value: joinId}}
		err := db.Collection("invitations").FindOne(c, filter).Decode(&invitation)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invite not found"})
				c.Abort()
				return
			} else {
				panic(err)
			}
		}

		c.Next()
	}
}