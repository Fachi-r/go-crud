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
	r.POST("/posts", controllers.CreatePosts)
	r.Run()
}
