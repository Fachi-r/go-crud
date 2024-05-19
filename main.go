package main

import (
	"github.com/fachi-r/go-crud/controllers"
	"github.com/fachi-r/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)

	r.POST("/posts", controllers.CreatePosts)
	r.PUT("/posts/:id", controllers.UpdatePost)

	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
