package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func GetBooksOnAuthorName(db *gorm.DB) {
	var books []model.Book
	db.Raw("    SELECT * FROM books join authors ON books.authorId = authors.authorId WHERE authors.name = 'George R. R. Martin'").Scan(&books)
	fmt.Println(books)
}
