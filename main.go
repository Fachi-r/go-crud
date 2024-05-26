package main

import (
	"github.com/fachi-r/go-crud/controllers"
	"github.com/fachi-r/go-crud/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	/*
		1. Setup router with Gin
	*/
	r := gin.Default()
	r.Use(cors.Default())

	/*
		2. Load Static assets (CSS files, JavaScript files) and HTML files
	*/
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("assets/html/*")

	/*
		3. Setup website routes
	*/
	// HOME PAGE
	r.GET("/", controllers.IndexPage)
	// Validate receipt number
	r.GET("/validate/receipts/:id", controllers.Validate)
	// Validate student loan number
	r.GET("/validate/students/:id", controllers.Validate)

	// SPECIFIC FORMS based on FORMID
	r.GET("/forms/:formID", controllers.GetForm)

	// Get All, or a specific student record
	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:loanNumber", controllers.GetStudent)
	// Create a student Record
	r.POST("/api/students", controllers.CreateStudent)

	// ADMIN PAGE
	r.GET("/admin", controllers.AdminPage)

	/*
		4. Run the Server
	*/
	r.Run()
}
