package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func UpdateAddressOfBob(db *gorm.DB) {
	var u model.User
	var p model.Profile
	db.Where("username = ?", "bob").First(&u)
	db.Table("profiles").Where("user_id = ?", u.Users_id).Update("address", "NewAddres665").Scan(&p)
	fmt.Println(p)

}
