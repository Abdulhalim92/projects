package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func SelectUsers(db *gorm.DB) {
	var u []model.User
	db.Find(&u)
	fmt.Println(u)
}
