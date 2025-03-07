package service

import (
	"defaultproject/repository"
	"defaultproject/serializer"
)

func BookStoreGetBook(title string) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	result, err := repo.GetBook(title)
	r = serializer.GeneralResponse("", result)
	return
}

func BookStoreBorrowBook(title string, borrower string) (r serializer.Response, err error) {
	var repo repository.BookStoreRepository
	count, err := repo.CheckBookAvailability(title)

	r = serializer.GeneralResponse("", result)
}
