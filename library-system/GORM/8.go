package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func GetBorrowedBooks(db *gorm.DB) {
	var books []model.Book
	db.Joins("JOIN borrow ON books.books_id = borrow.book_id").Where("borrow.return_date IS NULL").Find(&books)
	fmt.Println(books)
}
