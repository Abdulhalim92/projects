package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"projects/internal/model"
)

func main() {
	dsn := "host=217.11.185.181 user=humo password=humo dbname=go_lessons_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//1 Выбрать всех пользователей
	var u []model.User
	result1 := db.Find(&u)
	if result1.Error != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", u)

	//2 Выбрать все книги
	var books []model.Book
	result := db.Find(&books)
	if result.Error != nil {
		panic(err)
	}

	//3 Найти всех авторов

	var authors []model.Author
	result2 := db.Find(&authors)
	if result2.Error != nil {
		panic(err)
	}

	//4 Выбрать книги определенного автора по имени
	var books1 []model.Book
	authorName := "John Doe"
	result3 := db.Joins("JOIN authors ON authors.id = books.author_id").
		Where("authors.name = ?", authorName).
		Find(&books1)
	if result3.Error != nil {
		panic(err)
	}
}
