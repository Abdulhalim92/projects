package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func UpdateBiographyOfLeo(db *gorm.DB) {
	var x model.Author
	db.Table("authors").Where("authors.name = ?", "Leo Tolstoy").Update("biography", "NewBiographyforLeo")
	db.Raw("Select * From authors WHERE authors.name = 'Leo Tolstoy'").Scan(&x)
	fmt.Println(x)
}
