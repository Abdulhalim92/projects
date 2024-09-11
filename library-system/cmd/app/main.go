package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"projects/internal"
	"projects/internal/handler"
	"projects/internal/repository"
	"projects/internal/service"
)

func OperateThroughCL() {
	for {
		var command string

		fmt.Scan(&command)
		if command == "stop" {
			break
		}
	}
}

func main() {
	//cd library-system/cmd/app
	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		fmt.Errorf("Failed to connect to Database: %v\n", err)
	}

	mux := http.NewServeMux()

	// Инициализация книг
	BookRepo := repository.NewBookRepo(db)
	bookService := service.NewBookService(*BookRepo)

	bookHandler := handler.NewBookHandler(mux, bookService)
	bookHandler.InitRoutes()

	// Инициализация пользователей
	UserRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(*UserRepo)

	userHandler := handler.NewUserHandler(mux, userService)
	userHandler.InitRoutes()

	handler := internal.NewMyHandler(mux, *bookHandler, *userHandler)

	fmt.Printf("Server is starting...address: %s", ":8080\n")

	http.ListenAndServe("localhost:8080", handler)

	//// Инициализация пользователей
	//UserRepo := user.NewUserRepo(db)
	//userService := user.NewBookService(*UserRepo)
	//
	//// Имитируем создание книг
	//bookService.CreateBook(&model.Book{Title: "The Hobbit", AuthorID: 3})
	////bookService.CreateBook("1984", "George Orwell")
	//
	//// Имитируем создание пользователей
	//userService.CreateUser(&model.User{Username: "johndoe", Password: "password123"})
	////u2 := userService.CreateUser("janedoe", "securepassword")
	//
	//// Имитируем получение списка пользователей
	//listUsers, err := userService.ListUsers()
	//fmt.Println("UsersRepo in system:")
	//for _, u := range listUsers {
	//	fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.UserID, u.Username, u.Password)
	//}
	//
	////// Имитируем получение списка книг
	//listBooks, err := bookService.ListBooks()
	//
	//fmt.Println("BooksRepo in library:")
	//for _, b := range listBooks {
	//	fmt.Printf("ID: %d, Title: %s, Author: %d\n", b.BookID, b.Title, b.AuthorID)
	//}
	////fmt.Println(listBooks)
	////
	////// Имитируем обновление пользователя
	////updatedUser := userService.EditUser(u1.UserID, "johnsmith", "newpassword")
	////if updatedUser {
	////	fmt.Printf("User ID %d updated successfully\n", u1.UserID)
	////}
	////
	////// Имитируем обновление книг
	////updatedBook := bookService.EditBook(b1.BookID, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	////if updatedBook {
	////	b1 = *bookService.FindBook(b1.BookID)
	////	fmt.Printf("Book ID %d updated successfully: Title: %s, Author: %s\n", b1.BookID, b1.Title, b1.Author)
	////}
	////
	////// Имитируем удаление пользователя
	////deletedUser := userService.RemoveUser(u2.UserID)
	////if deletedUser {
	////	fmt.Printf("User ID %d deleted successfully\n", u2.UserID)
	////}
	////// Имитируем получение пользователя по ID
	////foundUser := userService.FindUser(u1.UserID)
	////if foundUser != nil {
	////	fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.UserID, foundUser.Username, foundUser.Password)
	////}
	////
	//// Имитируем получение книги по ID
	//foundBook, err := bookService.FindBook(1)
	//if err == nil {
	//	fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.BookID, foundBook.Title, foundBook.AuthorID)
	//} else {
	//	fmt.Printf("Couldn't find Book with given ID")
	//}
	//// Имитируем получение книги по author_id
	//booksByAuthor, err := bookService.FindBooksByAuthor(3)
	//fmt.Println("BooksRepo by author:")
	//for _, b := range booksByAuthor {
	//	fmt.Printf("ID: %d, Title: %s, Author: %d\n", b.BookID, b.Title, b.AuthorID)
	//}
	//
	//// Имитируем удаление книги
	//deletedBook, err := bookService.RemoveBook(3)
	//if err != nil {
	//	fmt.Printf("Book ID %d deleted successfully\n", deletedBook)
	//}

}

func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=muqaddas password=password dbname=library_db port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %v\n", err)
	}
	return db, nil
}
