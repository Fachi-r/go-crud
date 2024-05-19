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

func GetPosts(c *gin.Context) {
	// Get Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with Posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")
	// Get Posts
	var post models.Post
	initializers.DB.Find(&post, id)

	// Respond with Posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")
	// get post with ID
	var post models.Post
	initializers.DB.Find(&post, id)

	// Get Data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// update post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with Posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Delete Post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with status
	c.Status(200)
}
