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

	r := gin.Default()
	r.Use(cors.Default())
	// Load Static assets (CSS files, JavaScript files)
	r.Static("/assets", "./assets")
	// Load HTML files
	r.LoadHTMLGlob("assets/html/*")

	/* 2. Setup website routes*/
	// Home Page (Login with receipt)
	r.GET("/", controllers.IndexPage)

	// Admin Page routes
	r.GET("/admin", controllers.AdminPage)

	// After login, serve the form specified by formID
	r.GET("/forms/:id", controllers.GetForm)

	// Validate receipt number
	r.GET("/api/receipts/:receiptNumber", controllers.CheckReceipt)
	// Validate student loan number
	// Get a specific student record
	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:loanNumber", controllers.GetStudent)

	r.Run()
}
