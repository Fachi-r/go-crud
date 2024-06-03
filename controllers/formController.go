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
	formID := c.Param("id")

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

// Creates a student record with the fields
// provided through the request body
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

/*
Check if Student or Guardian with `:id` exists in the database

Returns JSON in the form:

	{ "exists": true || false }
*/
func Get(c *gin.Context) {
	id := c.Param("id")
	path := c.FullPath()

	var result *gorm.DB
	var data interface{}

	switch path {
	case "/api/students/:id":
		var Model models.Student
		// Check if a specific value exists in the students table
		result = database.DB.Where("loan_number = ?", id).First(&Model)
		data = Model

	case "/api/guardians/:id":
		var Model models.Guardian
		// Check if a specific value exists in the students table
		result = database.DB.Where("id = ?", id).First(&Model)
		data = Model
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to fetch record")
		return
	} else {
		// Value exists in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

/*
Return an array of all records in the Student or Guardian Database

Returns JSON in the form:

	{
		"data": [{...},{...}]
	}
*/
func GetAll(c *gin.Context) {
	path := c.FullPath()

	var result *gorm.DB
	var data interface{}

	switch path {
	case "/api/students":
		var Model []models.Student
		// Check if a specific value exists in the students table
		result = database.DB.Find(&Model)
		data = Model

	case "/api/guardians":
		var Model []models.Guardian
		// Check if a specific value exists in the students table
		result = database.DB.Find(&Model)
		data = Model
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to fetch record")
		return
	} else {
		// Value exists in the receipts table
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func Update(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")
	path := c.FullPath()
	// Get Data from request body
	var body struct {
		models.Student
		models.Guardian
	}
	c.Bind(&body)

	var result *gorm.DB
	var data interface{}

	switch path {
	case "/api/students/:id":
		var Model models.Student
		// Check if a specific value exists in the students table
		result = database.DB.Where("loan_number = ?", id).First(&Model)
		// update student fields
		// These are the only fields that are getting updated in the returning students' form
		database.DB.Model(&Model).Updates(models.Student{
			YearOfStudy:   body.YearOfStudy,
			Bank:          body.Bank,
			Branch:        body.Branch,
			AccountName:   body.AccountName,
			AccountNumber: body.AccountNumber,
		})

		data = Model

	case "/api/guardians/:id":
		var Model models.Guardian
		// get post with ID
		result = database.DB.Where("id = ?", id).First(&Model)

		// update guardian fields
		database.DB.Model(&Model).Updates(models.Guardian{
			GuardianName:  body.GuardianName,
			GuardianNRC:   body.GuardianNRC,
			Nationality:   body.Nationality,
			Gender:        body.Gender,
			Relationship:  body.Relationship,
			Address:       body.Address,
			Town:          body.Town,
			Province:      body.Province,
			PostalAddress: body.PostalAddress,
			Phone:         body.Phone,
			Email:         body.Email,
		})

		data = Model
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusInternalServerError)
		log.Fatal("Error: Failed to fetch record")
		return
	} else {
		// Return record
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func Delete(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	path := c.FullPath()
	switch path {
	case "/api/students/:id":
		// Delete Student
		database.DB.Delete(&models.Student{}, id)

	case "/api/guardians/:id":
		// Delete Guardian
		database.DB.Delete(&models.Guardian{}, id)

	case "/api/receipts/:id":
		// Delete Receipt
		database.DB.Delete(&models.Receipts{}, id)
	}

	// Respond with status
	c.Status(200)
}
