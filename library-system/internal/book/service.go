package book

import "projects/library-system/internal/model"

func CreateBook(title, author string) model.Book {
	return AddBook(title, author)
}

func ListBooks() []model.Book {
	return GetBooks()
}

func FindBook(id int) *model.Book {
	return GetBookByID(id)
}

func EditBook(id int, title, author string) bool {
	return UpdateBook(id, title, author)
}

func RemoveBook(id int) bool {
	return DeleteBook(id)
}
