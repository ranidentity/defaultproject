package response

import "time"

type GeneralRequest struct {
	EventId uint     `form:"event_id" json:"event_id"`
	Seats   []string `form:"seats" json:"seats"`
}

type BookStoreRequest struct {
	Title          string    `form:"title" json:"title"`
	NameOfBorrower string    `form:"name_of_borrower" json:"name_of_borrower"`
	BorrowDate     time.Time `form:"borrow_date" json:"borrow_date"`
	ReturnDate     time.Time `form:"return_date" json:"return_date"`
	LoanId         uint      `form:"loan_id" json:"loan_id"`
}
