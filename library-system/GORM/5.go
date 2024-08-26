package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func SelectProfilesWithEmail(db *gorm.DB) {
	var profiles []model.Profile
	db.Table("profiles").Where("profiles.email IS NOT NULL").Scan(&profiles)
	fmt.Println(profiles)
}
