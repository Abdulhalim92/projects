package main

import (
	"encoding/json"
	"fmt"
	"os"
	"projects/library-system/internal/book"
	"projects/library-system/internal/model"
	"projects/library-system/internal/user"
	// "projects/internal/book"
	// "projects/internal/model"
	// "projects/internal/user"
)

func main() {
	fmt.Println("Library System")

	// Инициализация книг
	//books := make(map[int]model.Book)
	//newBooks := book.NewBooks(books)
	jsonBooks := book.NewJSONBooks("books.json")
	bookService := book.NewService(*jsonBooks)

	// Инициализация пользователей
	users := make(map[int]model.User)
	newUsers := user.NewUsers(users)
	userService := user.NewService(*newUsers)

	// Имитируем создание книг
	b1 := bookService.CreateBook("The Hobbit", "J.R.R Tolkien")
	b2 := bookService.CreateBook("1984", "George Orwell")

	// Имитируем создание пользователей
	u1 := userService.CreateUser("johndoe", "password123")
	u2 := userService.CreateUser("janedoe", "securepassword")

	// Имитируем получение списка пользователей
	listUsers := userService.ListUsers()
	fmt.Println("Users in system:")
	for _, u := range listUsers {
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.ID, u.Username, u.Password)
	}

	// Имитируем получение списка книг
	listBooks := bookService.ListBooks()
	fmt.Println("Books in library:")
	for _, b := range listBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
	}

	// Имитируем обновление пользователя
	updatedUser := userService.EditUser(u1.ID, "johnsmith", "newpassword")
	if updatedUser {
		fmt.Printf("User ID %d updated successfully\n", u1.ID)
	}

	// Имитируем обновление книг
	updatedBook := bookService.EditBook(b1.ID, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	if updatedBook {
		fmt.Printf("Book ID %d updated successfully: Title: %s, Author: %s\n", b1.ID, b1.Title, b1.Author)
	}

	// Имитируем удаление пользователя
	deletedUser := userService.RemoveUser(u2.ID)
	if deletedUser {
		fmt.Printf("User ID %d deleted successfully\n", u2.ID)
	}

	// Имитируем удаление книги
	deletedBook := bookService.RemoveBook(b2.ID)
	if deletedBook {
		fmt.Printf("Book ID %d deleted successfully\n", b2.ID)
	}

	// Имитируем получение пользователя по ID
	foundUser := userService.FindUser(u1.ID)
	if foundUser != nil {
		fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.ID, foundUser.Username, foundUser.Password)
	}

	// Имитируем получение книги по ID
	foundBook := bookService.FindBook(b1.ID)
	if foundBook != nil {
		fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.ID, foundBook.Title, foundBook.Author)
	}
}

func WriteJSONToFile(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}
