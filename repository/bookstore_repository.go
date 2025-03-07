package repository

import (
	"defaultproject/model"
	"fmt"
)

type BookStoreRepository struct {
	model.Book
}

func (ref *BookStoreRepository) GetBook(title string) (result []model.Book, err error) {
	db := model.DB.Model(ref.Book)
	db.Where("title = ?", title)
	err = db.Find(&result).
		Error
	return
}

// check book availability
func (ref *BookStoreRepository) CheckBookAvailability(title string) (num int, err error) {
	var book []model.Book
	db := model.DB.Model(ref.Book)
	db.Where("title = ?", title)
	err = db.Find(&book).
		Error
	if len(book) > 1 {
		err = fmt.Errorf("'%s' is not precise enough ", title)
	}
	if book[0].AvailableCopies <= 0 {
		// err_msg = fmt.Sprintf("%s is not available now.", title)
		err = fmt.Errorf("book with title '%s' not found", title)
	}
	num = book[0].AvailableCopies
	return
}
