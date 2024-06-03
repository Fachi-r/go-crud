package database

import (
	"github.com/fachi-r/go-crud/models"
)

func Migrate() {

	// Push the database models to the database
	DB.AutoMigrate(&models.Student{}, &models.HELSB{}, &models.Receipts{}, &models.Guardian{})
	// Uploading the receipts to the database remotely
	receipts := []int{
		9848116,
		3482454,
		4950452,
		4754198,
		1224566,
		4123071,
		7524276,
		2775063,
		1801019,
		3871548,
	}

	for _, receipt := range receipts {
		newReceipt := models.Receipts{
			Number: uint(receipt),
		}
		DB.Create(&newReceipt)
	}
}
