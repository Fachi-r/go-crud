package models

/* Base level of the form.
Contains all fields in common between all the other forms
*/
type Form struct {
	Student
	Bank
	Guardian
}

type FirstYearForm struct {
	NRC           string `json:"nrc"`
	Name          string `json:"name"`
	YearOfStudy   string `json:"year_of_study"`
	Programme     string `json:"programme"`
	StudentNumber int64  `json:"student_number"`

	ReceiptNumber int64 `json:"receipt_number"`
	LoanNumber    int64 `json:"loan_number"`

	Bank          string `json:"bank"`
	Branch        string `json:"branch"`
	AccountName   string `json:"account_name"`
	AccountNumber int64  `json:"account_number"`

	Guardian     string `json:"guardian"`
	Gender       string `json:"gender"`
	Nationality  string `json:"nationality"`
	Relationship string `json:"relationship"`

	Address       string `json:"address"`
	Town          string `json:"town"`
	Province      string `json:"province"`
	PostalAddress string `json:"postal_address"`
	Phone         uint   `json:"phone"`
	Email         string `json:"email"`
}

type ReturningForm struct {
	Student
}

type Student struct {
	NRC int64 `json:"nrc" gorm:"primarykey"`

	Name          string `json:"name"`
	Programme     string `json:"programme"`
	YearOfStudy   string `json:"year_of_study"`
	StudentNumber int64  `json:"student_number"`

	ReceiptNumber int64 `json:"receipt_number"`
	LoanNumber    int64 `json:"loan_number"`
}

type Bank struct {
	BankName      string `json:"bank" gorm:"primarykey"`
	Branch        string `json:"branch"`
	AccountName   string `json:"account_name"`
	AccountNumber int64  `json:"account_number"`
}

type Guardian struct {
	GuardianName string `json:"guardian"`
	Relationship string `json:"relationship"`
	Gender       string `json:"gender"`
	Nationality  string `json:"nationality"`

	Address  string `json:"address"`
	Town     string `json:"town"`
	Province string `json:"province"`

	PostalAddress string `json:"postal_address"`
	Phone         uint   `json:"phone"`
	Email         string `json:"email"`
}
