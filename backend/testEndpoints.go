package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Frontend e2e testing uses this route
func resetDatabase(c *gin.Context) {
	resetDB()
	c.Status(http.StatusOK)
}
