package main

import (
	"fmt"
	"go-lesson/library-system/internal/book"
)

func main() {
	fmt.Println("Library System")

	// Иммитируем создание книг
	b1 := book.CreateBook("The Hobbit", "J.R.R Tolkien")
	b2 := book.CreateBook("1984", "George Orwell")

	// Иммитируем получение спискм книг
	books := book.ListBooks()
	fmt.Println("Books in library:")
	for _, b := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
	}

	// Имитируем обновление книги
	updated := book.EditBook(b1.ID, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	if updated {
		fmt.Printf("Book ID %d updated successfully\n", b1.ID)
	}

	// Имитируем удаление книги
	deleted := book.RemoveBook(b2.ID)
	if deleted {
		fmt.Printf("Book ID %d deleted successfully\n", b2.ID)
	}

	// Имитируем получение книги по ID
	foundBook := book.FindBook(b1.ID)
	if foundBook != nil {
		fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.ID, foundBook.Title, foundBook.Author)
	}
}
