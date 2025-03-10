package repository

import (
	"defaultproject/model"
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type BookStoreRepository struct {
	model.Book
	model.LoanDetail
}

func (ref *BookStoreRepository) GetBook(title string) (result []model.Book, err error) {
	db := model.DB.Model(ref.Book)
	if title != "" {
		db.Where("title = ?", title)
	}
	err = db.Find(&result).
		Error
	return
}

func (ref *BookStoreRepository) CheckBookAvailability(title string) (result model.Book, err error) {
	var book []model.Book
	db := model.DB.Model(ref.Book)
	db.Where("title = ?", title)
	err = db.Find(&book).
		Error
	switch {
	case len(book) > 1:
		err = fmt.Errorf("'%s' is not precise enough ", title)
	case len(book) == 0:
		err = fmt.Errorf("book with title '%s' not found", title)
	case book[0].AvailableCopies <= 0:
		err = fmt.Errorf("book with title '%s' is not available now", title)
	default:
		result = book[0]
	}
	return
}

func (ref *BookStoreRepository) LoanBook(book_id uint, new_count int, from time.Time, to time.Time, name_of_borrower string) (affected_row int, err error) {
	tx := model.DB.Begin()
	new := model.LoanDetail{
		NameOfBorrower: name_of_borrower,
		LoanDate:       from,
		ReturnDate:     to,
		BookId:         book_id,
	}
	if insert_error := tx.Debug().Table("loan_detail").Create(&new).Error; insert_error != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to initiate loan: %v", err)
	}
	result := tx.Model(ref.Book).
		Where("ID = ?", book_id).
		Where("available_copies > 0").
		Update("available_copies", new_count)
	if result.Error != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to update book's availability: %v", err)
	} else {
		affected_row = int(result.RowsAffected)
	}
	err = tx.Commit().Error
	return
}
func (ref *BookStoreRepository) ExtendLoan(loan_id uint) (affected_row int, err error) {
	tx := model.DB.Begin()
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", loan_id).
		First(&ref.LoanDetail).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to lock and fetch loan: %v", err)
	}
	ref.LoanDetail.ReturnDate = ref.LoanDetail.ReturnDate.AddDate(0, 0, 21)
	if err := tx.Save(&ref.LoanDetail).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to extend loan: %v", err)
	}
	affected_row = 1 // save doesnt trigger tx.RowAffected, so manually add 1
	err = tx.Commit().Error
	return
}

func (ref *BookStoreRepository) ReturnBook(loan_id uint) (affected_row int, err error) {
	tx := model.DB.Begin()
	// Lock the row for update
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", loan_id).
		First(&ref.LoanDetail).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to lock and fetch loan: %v", err)
	}
	currentTime := time.Now()
	ref.LoanDetail.BookReturnedOn = &currentTime
	if err := tx.Save(&ref.LoanDetail).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to update loan return book: %v", err)
	}
	book_id := ref.LoanDetail.BookId
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", book_id).
		First(&ref.Book).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to lock and fetch book: %v", err)
	}
	ref.Book.AvailableCopies = ref.Book.AvailableCopies + 1
	if err := tx.Save(&ref.Book).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to update book available copy: %v", err)
	}
	affected_row = 1
	err = tx.Commit().Error
	return
}
