package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		if !exists {
			log.Println("not authorized")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
