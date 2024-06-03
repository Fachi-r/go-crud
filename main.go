package main

import (
	"github.com/fachi-r/go-crud/controllers"
	// "github.com/fachi-r/go-crud/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	// database.LoadEnvVariables()
// 	// database.ConnectToDB()
// }

func main() {
	/*
		1. Setup router with Gin
	*/
	r := gin.Default()
	r.Use(cors.Default())

	/*
		2. Load Static assets (CSS files, JavaScript files, and HTML files)
	*/
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("assets/html/*")

	/*
		3. Setup website routes
	*/
	// HOME PAGE
	r.GET("/", controllers.IndexPage)
	r.GET("/login", controllers.IndexPage)
	// Validate receipts and student
	r.GET("/validate/receipts/:id", controllers.Validate)
	r.GET("/validate/students/:id", controllers.Validate)

	r.GET("/redirect/:id", controllers.Redirect)

	// Serve specific forms
	r.GET("/forms/:id", controllers.GetForm)
	r.GET("/forms/:id/docs", controllers.GetFormDocs)

	// CREATE record
	r.POST("/api/students", controllers.CreateStudent)
	r.POST("/api/forms/:id/docs", controllers.UploadFiles)

	// RETRIEVE (all, or a specific student and guardian records)
	r.GET("/api/students", controllers.GetAll)
	r.GET("/api/guardians", controllers.GetAll)
	r.GET("/api/students/:id", controllers.Get)
	r.GET("/api/guardians/:id", controllers.Get)

	// UPDATE records
	r.PUT("/api/students/:id", controllers.Update)
	r.PUT("/api/guardians/:id", controllers.Update)

	// DELETE records
	r.DELETE("/api/students/:id", controllers.Delete)
	r.DELETE("/api/guardians/:id", controllers.Delete)
	r.DELETE("/api/receipts/:id", controllers.Delete)

	// ADMIN PAGE
	r.GET("/admin", controllers.AdminPage)

	/*
		4. Run the Server
	*/
	r.Run()
}
