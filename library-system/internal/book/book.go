package book

import (
	"fmt"
	"projects/internal/model"
)

type Books struct {
	Books []model.Book
}

func NewBooks() *Books {
	books := make([]model.Book, 0)
	return &Books{Books: books}
}

var lastID int

func init() {
	lastID = 0
}

func (b *Books) AddBook(title, author string) model.Book {
	lastID++
	book := model.Book{
		ID:     lastID,
		Title:  title,
		Author: author,
	}

	b.Books = append(b.Books, book)

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *Books) GetBooks() []model.Book {
	return b.Books
}

func (b *Books) GetBookByID(id int) *model.Book {
	for _, book := range b.Books {
		if book.ID == id {
			fmt.Printf("Found book with id %d: %+v\n", id, book)
			return &book
		}
	}
	return nil
}

func (b *Books) GetBooksByAuthor(author string) []model.Book {
	var booksByAuthor []model.Book

	for _, book := range b.Books {
		if book.Author == author {
			booksByAuthor = append(booksByAuthor, book)
		}
	}
	return booksByAuthor
}

func (b *Books) UpdateBook(id int, title, author string) bool {
	for i, book := range b.Books {
		if book.ID == id {
			b.Books[i].Title = title
			b.Books[i].Author = author
			return true
		}
	}
	return false
}

func (b *Books) DeleteBook(id int) bool {
	for i, book := range b.Books {
		if book.ID == id {
			b.Books = append(b.Books[:i], b.Books[i+1:]...)
			return true
		}
	}
	return false
}
