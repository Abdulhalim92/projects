package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func ShowAllUsersWithoutProfiles(db *gorm.DB) {
	var users []model.User
	db.Joins("LEFT JOIN profiles ON users.users_id = profiles.user_id").Select("users_id", "username", "password").Where("user_id IS NULL").Find(&users)
	fmt.Println(users)
}
