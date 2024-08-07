package book

import (
	"fmt"
	"projects/internal/model"
)

type Books struct {
	BooksMap map[int]model.Book
	lastID   int
}

func NewBooks(books map[int]model.Book) *Books {
	return &Books{
		BooksMap: books,
		lastID:   0,
	}
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

	b.BooksMap[lastID] = book

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *Books) GetBooks() map[int]model.Book {
	return b.BooksMap
}

func (b *Books) GetBookByID(id int) *model.Book {
	for k, book := range b.BooksMap {
		if k == id {
			fmt.Printf("Found book with id %d: %+v\n", id, book)
			return &book
		}
	}
	return nil
}

func (b *Books) GetBooksByAuthor(author string) []model.Book {
	var booksByAuthor []model.Book

	for _, book := range b.BooksMap {
		if book.Author == author {
			booksByAuthor = append(booksByAuthor, book)
		}
	}
	return booksByAuthor
}

func (b *Books) UpdateBook(book model.Book) bool {
	for k, _ := range b.BooksMap {
		if k == book.ID {
			b.BooksMap[k] = book
			return true
		}
	}
	return false
}

func (b *Books) DeleteBook(id int) bool {
	for k, _ := range b.BooksMap {
		if k == id {
			delete(b.BooksMap, id)
			return true
		}
	}
	return false
}
