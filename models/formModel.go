package models

import "gorm.io/gorm"

// Receipts Table
type Receipts struct {
	Number uint `gorm:"primarykey"`
}

// HELSB Table
type HELSB struct {
	CEO string
}

type Student struct {
	NRC           uint
	Name          string
	Programme     string
	YearOfStudy   string `json:"year_of_study"`
	StudentNumber uint   `json:"student_number"`
	LoanNumber    uint   `json:"loan_number" gorm:"primarykey"`

	// Bank Details
	Bank          string
	Branch        string
	AccountName   string `json:"account_name"`
	AccountNumber uint   `json:"account_number"`

	// ID of guardian
	Guardian uint
}

type Guardian struct {
	gorm.Model
	GuardianName string `json:"guardian"`
	Relationship string
	Gender       string
	Nationality  string

	Address  string
	Town     string
	Province string

	PostalAddress string `json:"postal_address"`
	Phone         uint
	Email         string
}
