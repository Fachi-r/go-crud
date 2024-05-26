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
	LoanNumber    uint `gorm:"primarykey"`
	NRC           string
	Name          string
	Programme     string
	Degree        string
	School        string
	YearOfStudy   uint
	StudentNumber uint

	// Bank Details
	Bank          string
	Branch        string
	AccountName   string
	AccountNumber uint

	// ID of guardian
	Guardian uint
}

type Guardian struct {
	gorm.Model
	GuardianName string
	GuardianNRC  string
	Relationship string
	Gender       string
	Nationality  string

	Address  string
	Town     string
	Province string

	PostalAddress string
	Phone         string
	Email         string
}
