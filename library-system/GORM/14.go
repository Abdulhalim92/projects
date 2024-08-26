package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func SelectBooksAfterDate(db *gorm.DB) {
	var books []model.Book
	db.Joins("JOIN borrow ON books.books_id = borrow.book_id").Select("books_id", "title", "author_id").Where("borrow.return_date > ?", "2022-01-30").Find(&books)
	fmt.Println(books)
}
