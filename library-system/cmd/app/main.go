package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projects/internal/book"
	"projects/internal/model"
)

func main() {
	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		panic(err) //TODO
	}

	// Инициализация книг
	bookRepo := book.NewBookRepo(db)
	bookService := book.NewService(*bookRepo)

	//books, err := bookService.ListBooks()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, book := range books {
	//	fmt.Printf("%+v\n", book)
	//}

	//book, err := bookService.FindBooksByAuthor(1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", book)

	book := &model.Book{
		Title:    "Book of Death",
		AuthorID: 3,
	}

	books, err := bookService.CreateBook(book)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Books", books)

	//fmt.Printf("%+v", books)

	// Инициализация пользователей

	//users := make(map[int]model.User)
	//newUsers := user.NewUsers(users)
	//userService := user.NewService(*newUsers)
	//
	//for {
	//	fmt.Println("Enter command:")
	//	var command string
	//	fmt.Scanln(&command)
	//	if command == "exit" {
	//		return
	//	}
	//
	//	switch command {
	//	case "create-book":
	//		fmt.Println("Enter book title:")
	//		var title string
	//		fmt.Scanln(&title)
	//		fmt.Println("Enter book author:")
	//		var author string
	//		fmt.Scanln(&author)
	//		//bookService.CreateBook(title, author)
	//	case "list-books":
	//		//listBooks := bookService.ListBooks()
	//		fmt.Println("Books in library:")
	//		//for _, b := range listBooks {
	//		//	fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.BookId, b.Title, b.Author)
	//		//}
	//	case "create-user":
	//		fmt.Println("Enter username:")
	//		var username string
	//		fmt.Scanln(&username)
	//		fmt.Println("Enter password:")
	//		var password string
	//		fmt.Scanln(&password)
	//		userService.CreateUser(username, password)
	//	case "list-users":
	//		listUsers := userService.ListUsers()
	//		fmt.Println("Users in system:")
	//		for _, u := range listUsers {
	//			fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.UserID, u.Username, u.Password)
	//		}
	//	case "edit-book":
	//		fmt.Println("Enter book ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		fmt.Println("Enter new title:")
	//		var title string
	//		fmt.Scanln(&title)
	//		fmt.Println("Enter new author:")
	//		var author string
	//		fmt.Scanln(&author)
	//		//updatedBook := bookService.EditBook(id, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	//		//if updatedBook {
	//		//	fmt.Printf("Book ID %d updated successfully: Title: %s, Author: %s\n", id, title, author)
	//		//}
	//	case "edit-user":
	//		fmt.Println("Enter user ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		fmt.Println("Enter new username:")
	//		var username string
	//		fmt.Scanln(&username)
	//		fmt.Println("Enter new password:")
	//		var password string
	//		fmt.Scanln(&password)
	//		userService.EditUser(id, username, password)
	//	case "delete-book":
	//		fmt.Println("Enter book ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		//deletedBook := bookService.RemoveBook(id)
	//		//if deletedBook {
	//		//	fmt.Printf("Book ID %d deleted successfully\n", id)
	//		//}
	//	case "delete-user":
	//		fmt.Println("Enter user ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		userService.RemoveUser(id)
	//	case "get-book-by-id":
	//		fmt.Println("Enter book ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		//foundBook := bookService.FindBook(id)
	//		//if foundBook != nil {
	//		//	fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.BookId, foundBook.Title, foundBook.Author)
	//		//} else {
	//		//	fmt.Printf("Book with ID %d not found\n", id)
	//		//}
	//	case "get-user-by-id":
	//		fmt.Println("Enter user ID:")
	//		var id int
	//		fmt.Scanln(&id)
	//		foundUser := userService.FindUser(id)
	//		if foundUser != nil {
	//			fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.UserID, foundUser.Username, foundUser.Password)
	//		} else {
	//			fmt.Printf("User with ID %d not found\n", id)
	//		}
	//	}
	//}

}

func connectToDB() (*gorm.DB, error) {
	dsn := "host=217.11.185.181 user=mr_de password=humo dbname=go_lessons_db_mr_de port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
