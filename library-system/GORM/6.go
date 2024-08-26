package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func UpdatePasswordOfAlice(db *gorm.DB) {
	var b model.User
	db.Model(&model.User{}).Where("username = ?", "alice").Update("password", "NewPassword23")
	db.Model(model.User{}).Where("username = ?", "alice").Scan(&b)
	fmt.Println(b)
}
