package main

import (
	"fmt"
	"projects/internal/book"
	"projects/internal/model"
	"projects/internal/user"
)

func main() {
	fmt.Println("Library System")

	books := make(map[int]model.Book)

	newBooks := book.NewBooks(books)
	service := book.NewService(*newBooks)

	// Иммитируем создание книг
	b1 := service.CreateBook("The Hobbit", "J.R.R Tolkien")
	b2 := service.CreateBook("1984", "George Orwell")

	// Иммитируем получение спискм книг
	listBooks := service.ListBooks()
	fmt.Println("Books in library:")
	for _, b := range listBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
	}

	// Имитируем обновление книги
	updated := service.EditBook(b1.ID, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	if updated {
		fmt.Printf("Book ID %d updated successfully\n", b1.ID)
	}

	// Имитируем удаление книги
	deleted := service.RemoveBook(b2.ID)
	if deleted {
		fmt.Printf("Book ID %d deleted successfully\n", b2.ID)
	}

	// Имитируем получение книги по ID
	foundBook := service.FindBook(b1.ID)
	if foundBook != nil {
		fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.ID, foundBook.Title, foundBook.Author)
	}
	// Users
	m := make(map[int]model.User)
	users := user.NewUsers(m)
	s := user.CreateService(*users)
	s.CreateUser(1, "ffg", "34fv34v3")
	fmt.Println(s)
	fmt.Println(s.ListUsers())
	fmt.Println(s.ListUserById(1))
	fmt.Println(s.RemoveUserById(1))
	fmt.Println()
}
