package book

import "projects/internal/model"

type BookInterface interface {
	AddBook(title, author string) model.Book
	GetBooks() map[int]model.Book
	GetBookByID(id int) model.Book
	GetBooksByAuthor(author string) map[int]model.Book
	UpdateBook(id int, title, author string) bool
	DeleteBook(id int) bool
}
