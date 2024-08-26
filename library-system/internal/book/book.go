package book

import (
	"fmt"
	"projects/internal/model"
)

type Books struct {
	BooksMap map[int]model.Book
	LASTID   int
}

func NewBooks(books map[int]model.Book) *Books {
	return &Books{
		BooksMap: books,
		LASTID:   0,
	}
}

func (b *Books) AddBook(title string, author int) model.Book {
	b.LASTID++
	book := model.Book{
		Books_id:  b.LASTID,
		Title:     title,
		Author_id: author,
	}
	b.BooksMap[b.LASTID] = book

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author_id)

	return b.BooksMap[b.LASTID]
}

func (b *Books) GetBooks() []model.Book {
	books := make([]model.Book, 0)
	for _, value := range b.BooksMap {
		books = append(books, value)
	}
	return books
}

func (b *Books) GetBookByID(id int) *model.Book {
	value, ok := b.BooksMap[id]
	if !ok {
		fmt.Printf("Does't exist\n")
		return nil
	}
	return &value
}

func (b *Books) GetBooksByAuthor(author int) []model.Book {
	var booksByAuthor []model.Book
	for _, value := range b.BooksMap {
		if value.Author_id == author {
			booksByAuthor = append(booksByAuthor, value)
		}
	}
	return booksByAuthor
}

func (b *Books) UpdateBook(id int, title string, author int) bool {
	for key := range b.BooksMap {
		if key == id {
			b.BooksMap[key] = model.Book{Title: title, Author_id: author}
			return true
		}
	}
	return false
}

func (b *Books) DeleteBook(id int) bool {
	for key := range b.BooksMap {
		if key == id {
			delete(b.BooksMap, key)
			return true
		}
	}
	return false
}
