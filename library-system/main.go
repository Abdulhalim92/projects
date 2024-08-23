package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	user_id  int
	Username string
	Password string
}

func main() {
	dsn := "host=localhost user=humo password=humo dbname=Humo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	var u []User
	err = db.Find(&u).Error
	if err != nil {
		log.Fatal(err)
	}
	var us User
	err = db.Last(&us).Error
	if err != nil {
		log.Fatal(err)
	}
	for i, user := range u {
		fmt.Println(i, user)
	}

}
