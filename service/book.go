package service

import (
	"defaultproject/repository"
	"defaultproject/serializer"
	"fmt"
	"time"
)

func BookStoreGetBook(title string) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	result, err := repo.GetBook(title)
	r = serializer.GeneralResponse("", result)
	return
}

// TODO sql lock row when check availability to prevent double book
// Note:: fixed loan duration - 10 days
func BookStoreBorrowBook(title string, borrower string) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	var msg string
	from := time.Now().Truncate(time.Minute).UTC()
	to := time.Now().Truncate(time.Minute).AddDate(0, 0, 10).UTC()
	book, err := repo.CheckBookAvailability(title)
	if err == nil {
		affected_row, err := repo.LoanBook(book.ID, book.AvailableCopies-1, from, to, borrower)
		if affected_row == 1 && err == nil {
			msg = fmt.Sprintf("Please return before %s", to)
		}
	}
	if err != nil {
		msg = err.Error()
	}
	r = serializer.MessageResponse(msg)
	return
}

func BookStoreExtendLoan(loan_id uint) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	var msg string
	affected_row, err := repo.ExtendLoan(loan_id)
	if affected_row == 1 && err == nil {
		msg = "Successfully extended row for another 3 weeks"
	}
	r = serializer.MessageResponse(msg)
	return
}

func BookStoreReturn(loan_id uint) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	var msg string
	affected_row, err := repo.ReturnBook(loan_id)
	if affected_row == 1 && err == nil {
		msg = "Book returned"
	}
	r = serializer.MessageResponse(msg)
	return
}
