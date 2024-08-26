package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func CountBooks(db *gorm.DB) {
	var n int
	db.Model(&model.Book{}).Raw("Select count(*) FROM books").Scan(&n)
	fmt.Println(n)
}
