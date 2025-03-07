package model

import "time"

func (LoanDetail) TableName() string {
	return "loan_detail"
}

type LoanDetail struct {
	BaseModel
	NameOfBorrower string    `json:"name_of_borrower"`
	LoanDate       time.Time `json:"loan_date"`
	ReturnDate     time.Time `json:"return_date"`
}
