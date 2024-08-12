package book

import (
	"fmt"
	"projects/internal/model"
)

type Books struct {
	Books  map[int]model.Book
	LastID int
}

func NewBooks(books map[int]model.Book) *Books {
	return &Books{
		Books: books,
	}
}

//var lastID int
//
//func init() {
//	lastID = 0
//}

func (b *Books) AddBook(title, author string) model.Book {
	b.LastID++
	book := model.Book{
		ID:     b.LastID,
		Title:  title,
		Author: author,
	}

	b.Books[b.LastID] = book
	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *Books) GetBooks() []model.Book {
	//return b.Books
	var books []model.Book
	for _, book := range b.Books {
		books = append(books, book)
	}
	return books
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
	//for i, book := range b.Books {
	//	if book.ID == id {
	//		b.Books[i].Title = title
	//		b.Books[i].Author = author
	//		return true
	//	}
	//}
	//return false
	book, ok := b.Books[id]
	if !ok {
		return false
	}
	book.Title = title
	book.Author = author
	fmt.Printf("Book with title %s and author %s is updated\n", book.Title, book.Author)
	b.Books[id] = book
	return true
}

func (b *Books) DeleteBook(id int) bool {
	//for i, book := range b.Books {
	//	if book.ID == id {
	//		b.Books = append(b.Books[:i], b.Books[i+1:]...)
	//		return true
	//	}
	//}
	//return false
	_, ok := b.Books[id]
	if !ok {
		return false
	}
	delete(b.Books, id)
	return true

}
