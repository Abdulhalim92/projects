package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func SelectAuthors(db *gorm.DB) {
	var a []model.Author
	db.Find(&a)
	fmt.Println(a)
}
