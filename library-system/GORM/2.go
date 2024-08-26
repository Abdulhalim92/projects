package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func SelectBooks(db *gorm.DB) {
	var b []model.Book
	db.Find(&b)
	fmt.Println(b)
}
