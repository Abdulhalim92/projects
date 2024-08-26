package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func ShowAuthorOfWarAndPeace(db *gorm.DB) {
	var a model.Author
	db.InnerJoins("INNER JOIN books ON authors.authors_id = books.author_id").Select("authors_id", "name", "biography").Find(&a)
	fmt.Println(a)
}
