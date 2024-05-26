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
	c.HTML(http.StatusOK, "login.html", nil)
}

/*
Check if Receipt or Student with `:ID` exists in the database

Returns JSON in the form:

	{ "exists": true || false }
*/
func Validate(c *gin.Context) {
	// get type from url
	modelID := c.Param("id")
	path := c.FullPath()
	var result *gorm.DB

	switch path {
	case "/validate/receipts/:id":
		var receiptModel models.Receipts
		// Check if value exists in the receipts table
		result = database.DB.Where("Number = ?", modelID).First(&receiptModel)

	case "/validate/students/:id":
		var studentModel models.Student
		// Check if value exists in the students table
		result = database.DB.Where("loan_number = ?", modelID).First(&studentModel)
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Value does not exist in the table
		c.JSON(http.StatusNotFound, gin.H{
			"exists": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exists": true,
		})
	}
}

/*
Reruns the form specified by the `:formId`.

	Example usage:
	  fetch("./forms/first")
	  fetch("./forms/returning")
*/
func GetForm(c *gin.Context) {
	// Get FormID from url
	formID := c.Param("formID")

	// Serve Form with matching formID
	switch formID {
	case "first":
		c.HTML(http.StatusOK, "firstYearForm.html", nil)
	case "returning":
		c.HTML(http.StatusOK, "returningForm.html", nil)
	default:
		c.HTML(http.StatusNotFound, "notFound.html", nil)
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
	var studentModel models.Student
	// Check if a specific value exists in the students table
	result := database.DB.Where("loan_number = ?", loanNumber).First(&studentModel)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to fetch Student record")
		return
	} else {
		// Value exists in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"student": studentModel,
		})
	}
}

// Creates a student record with the fields passed in through the request
func CreateStudent(c *gin.Context) {
	// Get Data from POST request
	var data struct {
		models.Student
		models.Guardian
	}
	c.Bind(&data)

	// Create a new record with Guardian Model
	guardian := models.Guardian{
		GuardianName:  data.GuardianName,
		GuardianNRC:   data.GuardianNRC,
		Relationship:  data.Relationship,
		Gender:        data.Gender,
		Nationality:   data.Nationality,
		Address:       data.Address,
		Town:          data.Town,
		Province:      data.Province,
		PostalAddress: data.PostalAddress,
		Phone:         data.Phone,
		Email:         data.Email,
	}

	// Upload guardian first
	guardianResult := database.DB.Create(&guardian)

	if guardianResult.Error != nil {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to add Guardian record")
		return
	} else {
		// If no error, upload student record with guardian ID as foreign key
		student := models.Student{
			NRC:           data.NRC,
			Name:          data.Name,
			Programme:     data.Programme,
			Degree:        data.Degree,
			School:        data.School,
			YearOfStudy:   data.YearOfStudy,
			StudentNumber: data.StudentNumber,
			LoanNumber:    data.LoanNumber,
			// Bank Details
			Bank:          data.Bank,
			Branch:        data.Branch,
			AccountName:   data.AccountName,
			AccountNumber: data.AccountNumber,
			// Guardian
			Guardian: guardian.ID, //ID of guardian which was just inserted
		}
		studentWrite := database.DB.Create(&student)

		// Return the post
		if studentWrite.Error != nil {
			c.Status(http.StatusInternalServerError)
			log.Fatal("Error: Failed to add Student record")
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"student":  student,
				"guardian": guardian,
			})
		}
	}
}
