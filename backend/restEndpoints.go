package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func projectsHandler(c *gin.Context) {
	//DEBUG: Dummy data
	response := ProjectsResponse{
		Projects: []minimalProject{
			{
				Id:          uuid.New().String(),
				Title:       "Project 1",
				Description: "This is project 1",
			},
		},
	}

	log.Println(response)
	c.JSON(http.StatusOK, response)
}
