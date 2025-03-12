package service

import (
	"defaultproject/model"
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

// Note:: fixed loan duration - 10 days
func BookStoreBorrowBook(title string, borrower string) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	var msg string
	var data model.LoanDetail
	from := time.Now().Truncate(time.Minute).UTC()
	to := time.Now().Truncate(time.Minute).AddDate(0, 0, 10).UTC()
	book, err := repo.CheckBookAvailability(title)
	// Sometimes book count > 1, locking it throughout the whole process will prevent other people from accessing the row and not able to rent
	if err == nil {
		data, err = repo.LoanBook(book.ID, book.AvailableCopies-1, from, to, borrower)
		if err == nil {
			msg = fmt.Sprintf("Please return before %s", to)
		}
	}
	if err != nil {
		msg = err.Error()
	}
	if data != (model.LoanDetail{}) {
		r = serializer.GeneralResponse(msg, data)
	} else {
		r = serializer.MessageResponse(msg)
	}
	return
}

// TODO return error message
func BookStoreExtendLoan(loan_id uint) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	var msg string
	affected_row, err := repo.ExtendLoan(loan_id)
	if affected_row == 1 && err == nil {
		msg = "Successfully extended loan for another 3 weeks"
	} else {
		msg = err.Error()
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
	} else {
		msg = err.Error()
	}
	r = serializer.MessageResponse(msg)
	return
}
