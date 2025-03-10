package response

import "time"

type BookStoreRequest struct {
	Title          string    `form:"title" json:"title"`
	NameOfBorrower string    `form:"name_of_borrower" json:"name_of_borrower"`
	BorrowDate     time.Time `form:"borrow_date" json:"borrow_date"`
	ReturnDate     time.Time `form:"return_date" json:"return_date"`
	LoanId         uint      `form:"loan_id" json:"loan_id"`
}

type BookStoreBorrowBookRequest struct {
	Title          string `form:"title" json:"title" binding:"required"`
	NameOfBorrower string `form:"name_of_borrower" json:"name_of_borrower" binding:"required"`
}

type BookStoreExtendLoanRequest struct {
	LoanId uint `form:"loan_id" json:"loan_id" binding:"required"`
}
