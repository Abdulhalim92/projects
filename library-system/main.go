package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"projects/internal/model"
)

func main() {

	dsn := "host=localhost user=muqaddas password=password dbname=library_db port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	//1-- Выбрать всех пользователей.
	var u []model.User
	err = db.Find(&u).Error
	if err != nil {
		panic(err)
	}
	log.Println(u)

	//2 -- Выбрать все книги.
	var books []model.Book
	err = db.Find(&books).Error
	if err != nil {
		panic(err)
	}
	log.Println(books)

	//3-- Найти всех авторов.
	var authors []model.Author
	err = db.Find(&authors).Error
	if err != nil {
		panic(err)
	}
	log.Println(authors)

	//4-- Выбрать книги определенного автора по имени.
	db.Raw("SELECT * FROM books WHERE author_id = (SELECT authors.author_id FROM authors WHERE name='J. K. Rowling')").Scan(&books)
	//db.Table("books").Select("title").
	log.Println(books)
}
