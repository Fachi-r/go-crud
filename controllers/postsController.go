package controllers

import (
	"github.com/fachi-r/go-crud/initializers"
	"github.com/fachi-r/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {
	// Get Data from POST request
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a new post
	post := models.Post{Title: body.Title, Body: body.Body}

	// Return the post
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
