package book

import (
	"projects/internal/model"
	"testing"
)

func TestCreateBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := NewBooks(books)
	s := NewService(*b)

	book := s.CreateBook("The Hobbit", "J.R.R. Tolkien")
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

func TestListBooks(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := NewBooks(books)
	s := NewService(*b)

	s.CreateBook("The Hobbit", "J.R.R. Tolkien")
	s.CreateBook("1984", "George Orwell")

	booksList := s.ListBooks()
	if len(booksList) != 2 {
		t.Errorf("Expected 2 books, got %d", len(booksList))
	}
}

func TestFindBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := NewBooks(books)
	s := NewService(*b)

	s.CreateBook("The Hobbit", "J.R.R. Tolkien")
	book := s.FindBook(1)
	if book == nil {
		t.Errorf("Expected to find book with ID 1")
	}
	if book.Title != "The Hobbit" {
		t.Errorf("Expected book title to be 'The Hobbit', got '%s'", book.Title)
	}
}

func TestEditBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := NewBooks(books)
	s := NewService(*b)

	s.CreateBook("The Hobbit", "J.R.R. Tolkien")
	updated := s.EditBook(1, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	if !updated {
		t.Errorf("Expected book to be updated")
	}

	book := s.FindBook(1)
	if book.Title != "The Hobbit: An Unexpected Journey" {
		t.Errorf("Expected book title to be 'The Hobbit: An Unexpected Journey', got '%s'", book.Title)
	}
}

func TestRemoveBook(t *testing.T) {
	books := make([]model.Book, 0) // Reset the books slice before each test
	lastID = 0                     // Reset the lastID before each test

	b := NewBooks(books)
	s := NewService(*b)

	s.CreateBook("The Hobbit", "J.R.R. Tolkien")
	deleted := s.RemoveBook(1)
	if !deleted {
		t.Errorf("Expected book to be deleted")
	}

	book := s.FindBook(1)
	if book != nil {
		t.Errorf("Expected book with ID 1 to be deleted")
	}
}
