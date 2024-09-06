package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"library-system/internal/book"
	"library-system/internal/user"
	"net/http"
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

	mux := http.NewServeMux()
	bookHandler := book.NewBookHandler(mux, bookService)
	bookHandler.InitRoutes()

	fmt.Printf("Server is starting... address: %v", ":8080\n")
	err = http.ListenAndServe("localhost:8080", bookHandler)
	if err != nil {
		panic(err)
	}

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
			book, err := bookService.CreateBook(title, authorID)
			if err != nil {
				fmt.Printf("Failed to create book: %v\n", err)
			}
			fmt.Printf("Book created: ID: %d, Title: %s, AuthorID: %d\n", book.BookId, book.Title, book.AuthorId)
		case "list-books":
			listBooks, err := bookService.ListBooks()
			if err != nil {
				fmt.Println(err)
			}

			for _, b := range listBooks {
				fmt.Printf("ID: %d, Title: %s, AuthorID: %d\n", b.BookId, b.Title, b.AuthorId)
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
			listUsers, err := userService.ListUsers()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Users in system:")
			for _, u := range listUsers {
				fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.UserId, u.Username, u.Password)
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
			err := bookService.EditBook(id, title, author)
			if err != nil {
				fmt.Printf("Book ID %d updated successfully: Title: %s, AuthorID: %s\n", id, title, author)
			} else {
				fmt.Println(err)
			}

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
			err := userService.EditUser(id, username, password)
			if err != nil {
				fmt.Printf("User ID %d updated successfully: username: %s, password: %s\n", id, username, password)
			} else {
				fmt.Println(err)
			}
		case "delete-book":
			fmt.Println("Enter book ID:")
			var id int
			fmt.Scanln(&id)
			if ok := bookService.RemoveBook(id); ok {
				fmt.Printf("Book ID %d deleted successfully\n", id)

			} else {
				fmt.Println("Error occured when deleting book")
			}

		case "delete-user":
			fmt.Println("Enter user ID:")
			var id int
			fmt.Scanln(&id)
			if ok := userService.RemoveUser(id); ok {
				fmt.Printf("User ID %d deleted successfully\n", id)

			} else {
				fmt.Println("Error occured when deleting user")
			}
		case "get-book-by-id":
			fmt.Println("Enter book ID:")
			var id int
			fmt.Scanln(&id)
			foundBook, err := bookService.FindBook(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(foundBook)
			}

		case "get-user-by-id":
			fmt.Println("Enter user ID:")
			var id int
			fmt.Scanln(&id)
			foundUser, err := userService.FindUser(id)
			if err != nil {
				fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.UserId, foundUser.Username, foundUser.Password)
			} else {
				fmt.Println(err)
			}
		}
	}

}

// connectToDB connects to the PostgreSQL database using the provided DSN.
//
// It returns a pointer to a gorm.DB object and an error if any occurred.
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=saiddis password=__1dIslo__ dbname=library port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
