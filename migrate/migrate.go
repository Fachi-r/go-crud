package main

import (
	"github.com/fachi-r/go-crud/initializers"
	"github.com/fachi-r/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
