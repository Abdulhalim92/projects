package book

import "projects/library-system/internal/model"

var books []model.Book
var lastID int

func init() {
	books = make([]model.Book, 0)
	lastID = 0
}

func AddBook(title, author string) model.Book {
	lastID++
	book := model.Book{
		ID:     lastID,
		Title:  title,
		Author: author,
	}

	books = append(books, book)

	return book
}

func GetBooks() []model.Book {
	return books
}

func GetBookByID(id int) *model.Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func UpdateBook(id int, title, author string) bool {
	for i, book := range books {
		if book.ID == id {
			books[i].Title = title
			books[i].Author = author
			return true
		}
	}
	return false
}

func DeleteBook(id int) bool {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}