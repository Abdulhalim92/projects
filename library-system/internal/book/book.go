package book

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBooks(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (b *BookRepository) AddBook(title string, author int) *model.Book {
	b.LASTID++
	book := model.Book{
		BooksId:  b.LASTID,
		Title:    title,
		AuthorId: author,
	}
	b.BooksMap[b.LASTID] = book

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.AuthorId)

	return b.BooksMap[b.LASTID]
}

func (b *BookRepository) GetBooks() []model.Book {
	books := make([]model.Book, 0)
	for _, value := range b.BooksMap {
		books = append(books, value)
	}
	return books
}

func (b *BookRepository) GetBookByID(id int) *model.Book {
	value, ok := b.BooksMap[id]
	if !ok {
		fmt.Printf("Does't exist\n")
		return nil
	}
	return &value
}

func (b *BookRepository) GetBooksByAuthor(author int) []model.Book {
	var booksByAuthor []model.Book
	for _, value := range b.BooksMap {
		if value.Author_id == author {
			booksByAuthor = append(booksByAuthor, value)
		}
	}
	return booksByAuthor
}

func (b *BookRepository) UpdateBook(id int, title string, author int) bool {
	for key := range b.BooksMap {
		if key == id {
			b.BooksMap[key] = model.Book{Title: title, AuthorId: author}
			return true
		}
	}
	return false
}

func (b *BookRepository) DeleteBook(id int) bool {
	for key := range b.BooksMap {
		if key == id {
			delete(b.BooksMap, key)
			return true
		}
	}
	return false
}
