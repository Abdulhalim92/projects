package main

import (
	"fmt"
	"projects/internal/book"
	"projects/internal/model"
	"projects/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		panic(err) // TODO
	}

	// Инициализация книг
	bookRepo := book.NewBookRepo(db)
	bookService := book.NewService(*bookRepo)

	// Инициализация пользователей
	userRepo := user.NewUserRepo(db)
	userService := user.NewService(*userRepo)

	for {
		fmt.Println("Enter command:")
		var command string
		fmt.Scanln(&command)
		if command == "exit" {
			return
		}

		switch command {
		case "create-book":
			fmt.Println("Enter book title:")
			var title string
			fmt.Scanln(&title)
			fmt.Println("Enter book author_id:")
			var authorID int
			fmt.Scanln(&authorID)
			book := &model.Book{Title: title, AuthorID: authorID}
			book, err := bookService.CreateBook(book)
			if err != nil {
				fmt.Printf("Failed to create book: %v\n", err)
			}
			fmt.Printf("Book created: ID: %d, Title: %s, AuthorID: %d\n", book.BookId, book.Title, book.AuthorID)
		case "list-books":
			listBooks, err := bookService.ListBooks()
			if err != nil {
				fmt.Printf("Failed to list books: %v\n", err)
			}
			fmt.Println("Books in library:")
			for _, b := range listBooks {
				fmt.Printf("ID: %d, Title: %s, AuthorID: %s\n", b.BookId, b.Title, b.AuthorID)
			}
		case "create-user":
			fmt.Println("Enter username:")
			var username string
			fmt.Scanln(&username)
			fmt.Println("Enter password:")
			var password string
			fmt.Scanln(&password)
			userService.CreateUser(username, password)
		case "list-users":
			listUsers := userService.ListUsers()
			fmt.Println("Users in system:")
			for _, u := range listUsers {
				fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.UserID, u.Username, u.Password)
			}
		case "edit-book":
			fmt.Println("Enter book ID:")
			var id int
			fmt.Scanln(&id)
			fmt.Println("Enter new title:")
			var title string
			fmt.Scanln(&title)
			fmt.Println("Enter new author:")
			var author string
			fmt.Scanln(&author)
			book := &model.Book{BookId: id, Title: title, AuthorID: author}
			book, err := bookService.EditBook(book)
			if err != nil {
				fmt.Printf("Failed to edit book: %v\n", err)
			}
			fmt.Printf("Book ID %d updated successfully: Title: %s, AuthorID: %s\n", id, title, author)
		case "edit-user":
			fmt.Println("Enter user ID:")
			var id int
			fmt.Scanln(&id)
			fmt.Println("Enter new username:")
			var username string
			fmt.Scanln(&username)
			fmt.Println("Enter new password:")
			var password string
			fmt.Scanln(&password)
			userService.EditUser(id, username, password)
		case "delete-book":
			fmt.Println("Enter book ID:")
			var id int
			fmt.Scanln(&id)
			id, err := bookService.RemoveBook(id)
			if err != nil {
				fmt.Printf("Failed to delete book: %v\n", err)
			}
			fmt.Printf("Book ID %d deleted successfully\n", id)
		case "delete-user":
			fmt.Println("Enter user ID:")
			var id int
			fmt.Scanln(&id)
			userService.RemoveUser(id)
		case "get-book-by-id":
			fmt.Println("Enter book ID:")
			var id int
			fmt.Scanln(&id)
			foundBook, err := bookService.FindBook(id)
			if err != nil {
				fmt.Printf("Failed to get book: %v\n", err)
			}
			if foundBook != nil {
				fmt.Printf("Found Book - ID: %d, Title: %s, AuthorID: %s\n", foundBook.BookId, foundBook.Title, foundBook.AuthorID)
			} else {
				fmt.Printf("Book with ID %d not found\n", id)
			}
		case "get-user-by-id":
			fmt.Println("Enter user ID:")
			var id int
			fmt.Scanln(&id)
			foundUser := userService.FindUser(id)
			if foundUser != nil {
				fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.UserID, foundUser.Username, foundUser.Password)
			} else {
				fmt.Printf("User with ID %d not found\n", id)
			}
		}
	}

}

// connectToDB connects to the PostgreSQL database using the provided DSN.
//
// It returns a pointer to a gorm.DB object and an error if any occurred.
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
