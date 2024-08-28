package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"projects/internal/BookDataBase"
)

func main() {
	db, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	BookRep := BookDataBase.NewBookRepository(db)
	ser := BookDataBase.NewService(BookRep)
	books, err := ser.ListBooks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
	books2, err := ser.FindBookByAuthor(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books2)
	//userNew := model.User{Username: "NewUser", Password: "Newusername"}

	/*for {
		var operation string
		fmt.Println("Choose operation:")
		fmt.Scan(&operation)
		if operation == "exit" {
			return
		}
		switch operation {
		case "getbooks":
			books := s.ListBooks()
			for _, book := range books {
				fmt.Println(book)
			}
		case "getbookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			book := s.FindBook(id)
			fmt.Println(book)
		case "deletebookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			b := s.RemoveBook(id)
			if !b {
				fmt.Println("Error while deleting")
			}
			fmt.Println("Book deleted")
		case "changebookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			var author int
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			var title string
			fmt.Println("Enter title:")
			fmt.Scan(&title)
			b := s.EditBook(id, title, author)
			if !b {
				fmt.Println("Error while editing a book")
			}
			fmt.Println("Book edited")
		case "getbooksbyauthor":
			var author int
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			books := s.FindBooksByAuthor(author)
			for _, book := range books {
				fmt.Println(book)
			}
		case "addbook":
			var model.M
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			fmt.Println("Enter title:")
			fmt.Scan(&title)
			s.CreateBook()
		}
	} */
}
func ConnectToDb() (*gorm.DB, error) {
	dsn := "host=localhost user=humo password=humo dbname=Humo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to db: %v", err)
		return nil, err
	}
	return db, nil
}
