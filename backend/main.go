package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ws", wsEndpoint)

	router.GET("/projects", projectsHandler)

	router.Run(":8080")
}
