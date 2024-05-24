package main

import (
	"github.com/fachi-r/go-crud/database"
	"github.com/fachi-r/go-crud/models"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Student{}, &models.HELSB{}, &models.Receipts{})
}
