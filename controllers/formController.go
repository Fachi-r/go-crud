package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/fachi-r/go-crud/database"
	"github.com/fachi-r/go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Return Index page
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Return Admin page
func AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

/*
Check if Receipt with `:receiptNumber` exists in the database

Returns JSON in the form:

	{ "exists": true || false }
*/
func CheckReceipt(c *gin.Context) {
	// Get receipt from url
	receiptNumber := c.Param("receiptNumber")
	var receiptModel models.Receipts
	// Check if a specific value exists in the receipts table
	result := database.DB.Where("Number = ?", receiptNumber).First(&receiptModel)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Value does not exist in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"exists": false,
		})
	} else {
		// Value exists in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"exists": true,
		})
	}
}

/*
Check if Student with `:loanNumber` exists in the database

Returns JSON in the form:

	{ "exists": true || false }
*/
func GetStudent(c *gin.Context) {
	// Get student loan number from url
	loanNumber := c.Param("loanNumber")
	var student models.Student
	// Check if a specific value exists in the students table
	result := database.DB.Where("LoanNumber = ?", loanNumber).First(&student)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Value does not exist in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"exists": false,
		})
	} else {
		// Value exists in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"exists":  true,
			"student": student,
		})
	}
}

// Used on the admin dashboard to display all students
func GetAllStudents(c *gin.Context) {
	var student models.Student
	database.DB.First(&student)
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

/*
Reruns the form specified by the formId.

	Example usage:
	  fetch("./forms/first")
	  fetch("./forms/returning")
*/
func GetForm(c *gin.Context) {
	// Get FormID from url
	formID := c.Param("formID")

	// Get Form which matches the formID from database
	switch formID {
	case "first":
		c.HTML(http.StatusOK, "firstYearForm.html", nil)
	case "returning":
		c.HTML(http.StatusOK, "returningForm.html", nil)
	default:
		c.HTML(http.StatusNotFound, "notFound.html", nil)
	}

}

// Creates a student record with the fields passed in through the request
func CreateStudent(c *gin.Context) {
	// Get Data from POST request
	var data struct {
		student  models.Student
		guardian models.Guardian
	}
	c.Bind(&data)

	// Create a new post
	student := models.Student{
		NRC:           data.student.NRC,
		Name:          data.student.Name,
		Programme:     data.student.Programme,
		YearOfStudy:   data.student.YearOfStudy,
		StudentNumber: data.student.StudentNumber,
		LoanNumber:    data.student.LoanNumber,
		// Bank Details
		Bank:          data.student.Bank,
		Branch:        data.student.Branch,
		AccountName:   data.student.AccountName,
		AccountNumber: data.student.AccountNumber,
	}
	guardian := models.Guardian{
		GuardianName:  data.guardian.GuardianName,
		Relationship:  data.guardian.Relationship,
		Gender:        data.guardian.Gender,
		Nationality:   data.guardian.Nationality,
		Address:       data.guardian.Address,
		Town:          data.guardian.Town,
		Province:      data.guardian.Province,
		PostalAddress: data.guardian.PostalAddress,
		Phone:         data.guardian.Phone,
		Email:         data.guardian.Email,
	}
	guardianResult := database.DB.Create(&guardian)
	studentWrite := database.DB.Create(&student)

	// Return the post
	if studentWrite.Error != nil {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to add Student record")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
	}
	// Return the post
	if guardianResult.Error != nil {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to add Guardian record")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"guardian": guardian,
		})
	}
}
