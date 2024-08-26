package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func ShowProfileOfAllUsers(db *gorm.DB) {
	var profiles []model.Profile
	db.Joins("JOIN users ON profiles.user_id = users.users_id").Find(&profiles)
	fmt.Println(profiles)
}
