package controllers

import (
	"log"
	"net/http"

	"github.com/fachi-r/go-crud/database"
	"github.com/fachi-r/go-crud/models"
	"github.com/gin-gonic/gin"
)

// Return Admin page
func AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

// Return all students
func GetAllStudents(c *gin.Context) {
	var student models.Student
	result := database.DB.First(&student)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to fetch Student records")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
	}
}
