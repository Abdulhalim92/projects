package book

import (
	"fmt"
	"projects/internal/model"
)

type Books struct {
	BooksMap map[int]model.Book
	LastID   int
}

func NewBooks(books map[int]model.Book) *Books {
	return &Books{
		BooksMap: books,
		LastID:   0,
	}
}

func (b *Books) AddBook(title, author string) model.Book {
	b.LastID++
	book := model.Book{
		ID:     b.LastID,
		Title:  title,
		Author: author,
	}

	b.BooksMap[book.ID] = book

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *Books) GetBooks() map[int]model.Book {
	return b.BooksMap
}

func (b *Books) GetBookByID(id int) *model.Book {
	book, exists := b.BooksMap[id]
	if !exists {
		fmt.Printf("Book with id %d not found\n", id)
		return nil
	}
	return &book
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
	_, exists := b.BooksMap[b.LastID]

	if !exists {
		fmt.Printf("Book with id %d not found\n", book.ID)
		return false
	}

	fmt.Printf("Book with id %d updated: Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	b.BooksMap[b.LastID] = book

	return true
}

func (b *Books) DeleteBook(id int) bool {
	_, exists := b.BooksMap[id]
	if !exists {
		fmt.Printf("Book with id %d not found\n", id)
		return false
	}

	delete(b.BooksMap, id)

	return true
}
