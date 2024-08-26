package book

import "projects/internal/model"

type BookInterface interface {
	AddBook(title string, author int) model.Book
	GetBooks() map[int]model.Book
	GetBookByID(id int) model.Book
	GetBooksByAuthor(author int) map[int]model.Book
	UpdateBook(id int, title string, author int) bool
	DeleteBook(id int) bool
}
