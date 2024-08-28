package main

import (
	"fmt"
	"gorm.io/gorm"
)

func ShowAmountOfBooksOfEachAuthor(db *gorm.DB) {
	type x struct {
		name   string
		amount int
	}
	var n x
	db.Raw("SELECT authors.name, count(*) AS Author_name FROM books JOIN authors ON books.authorId = authors.authorId  GROUP BY authorId;").Scan(&n)
	fmt.Println(n)

}
