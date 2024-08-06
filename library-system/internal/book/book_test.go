package book

import (
	"projects/internal/model"
	"testing"
)

func TestAddBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := Books{
		Books: books,
	}

	book := b.AddBook("The Hobbit", "J.R.R. Tolkien")
	if book.ID != 1 {
		t.Errorf("Expected book ID to be 1, got %d", book.ID)
	}
	if book.Title != "The Hobbit" {
		t.Errorf("Expected book title to be 'The Hobbit', got '%s'", book.Title)
	}
	if book.Author != "J.R.R. Tolkien" {
		t.Errorf("Expected book author to be 'J.R.R. Tolkien', got '%s'", book.Author)
	}
}

func TestGetBooks(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := Books{
		Books: books,
	}

	b.AddBook("The Hobbit", "J.R.R. Tolkien")
	b.AddBook("1984", "George Orwell")

	booksList := b.GetBooks()
	if len(booksList) != 2 {
		t.Errorf("Expected 2 books, got %d", len(booksList))
	}
}

func TestGetBookByID(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := Books{
		Books: books,
	}

	b.AddBook("The Hobbit", "J.R.R. Tolkien")
	book := b.GetBookByID(1)
	if book == nil {
		t.Errorf("Expected to find book with ID 1")
	}
	if book.Title != "The Hobbit" {
		t.Errorf("Expected book title to be 'The Hobbit', got '%s'", book.Title)
	}
}

func TestUpdateBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := Books{
		Books: books,
	}

	b.AddBook("The Hobbit", "J.R.R. Tolkien")
	updated := b.UpdateBook(1, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	if !updated {
		t.Errorf("Expected book to be updated")
	}

	book := b.GetBookByID(1)
	if book.Title != "The Hobbit: An Unexpected Journey" {
		t.Errorf("Expected book title to be 'The Hobbit: An Unexpected Journey', got '%s'", book.Title)
	}
}

func TestDeleteBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := Books{
		Books: books,
	}

	b.AddBook("The Hobbit", "J.R.R. Tolkien")
	deleted := b.DeleteBook(1)
	if !deleted {
		t.Errorf("Expected book to be deleted")
	}

	book := b.GetBookByID(1)
	if book != nil {
		t.Errorf("Expected book with ID 1 to be deleted")
	}
}
