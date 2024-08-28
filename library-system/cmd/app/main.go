package main

import (
	"fmt"
	"log"

	"library-system/internal/book"
	"library-system/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=saiddis password=__1dIslo_ dbname=library port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// homework(db)
	bookRepo := book.NewBookRepository(db)
	// bookService := book.NewService(*bookRepo)

	if err != nil {
		fmt.Println(err)
		return
	}

	books, err := bookRepo.GetBooks()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range books {
		fmt.Println(v)
	}
}

func homework(db *gorm.DB) {

	fmt.Println("Ex1")
	var u []model.User
	err := db.Find(&u).Error
	if err != nil {
		panic(err)
	}

	for _, v := range u {
		fmt.Println(v)
	}

	fmt.Println("---Ex2---")
	var b []model.Book
	err = db.Find(&b).Error
	if err != nil {
		panic(err)
	}
	for _, book := range b {
		fmt.Println(book)
	}
	fmt.Println("---Ex3---")
	var a []model.Author
	err = db.Find(&a).Error
	if err != nil {
		panic(err)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("---Ex4---")
	var leoTolstoyTitles []string

	db.Table("books").
		Select("books").
		Joins("JOIN authors ON books.author_id = authors.id").
		Where("authors.name = ?", "Leo Tolstoy").Scan(&leoTolstoyTitles)

	for _, book := range leoTolstoyTitles {
		fmt.Println(book)
	}

	fmt.Println("---Ex4---")
	var p []model.Profiles
	err = db.Find(&p).Error
	if err != nil {
		panic(err)
	}

	var usersWithEmail []string

	db.Table("users").
		Select("users").
		Joins("JOIN profiles ON profiles.user_id = users.id").
		Where("profiles.email IS NOT NULL").Scan(&usersWithEmail)
	for _, user := range usersWithEmail {
		fmt.Println(user)
	}

	fmt.Println("---Ex5---")
	var alice string
	db.Save(&model.User{UserID: 1, Username: "alice", Password: "newpassword"})
	db.Table("users").
		Select("users").
		Where("username = ?", "alice").Scan(&alice)
	fmt.Println(alice)
}
