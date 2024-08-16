package book

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

const BooksFile = "books.json"

type JSONBooks struct {
	filename string
}

func NewJSONBooks(filename string) *JSONBooks {
	return &JSONBooks{
		filename: filename,
	}
}

func (b *JSONBooks) loadBooks() (map[int]model.Book, error) {
	books := make(map[int]model.Book)
	err := utils.ReadJSONFromFile(b.filename, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *JSONBooks) saveBooks(books map[int]model.Book) error {
	return utils.WriteJSONToFile(b.filename, &books)
}

func (b *JSONBooks) AddBook(title, author string) model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return model.Book{}
	}
	lastID := 0

	for id := range books {
		if id >= lastID {
			lastID = id
		}
	}
	lastID++

	book := model.Book{
		ID:     lastID,
		Title:  title,
		Author: author,
	}

	books[lastID] = book

	err = b.saveBooks(books)
	if err != nil {
		fmt.Printf("Failed to save books: %v\n", err)
		return model.Book{}
	}

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *JSONBooks) GetBooks() map[int]model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	return books
}

func (b *JSONBooks) GetBookByID(id int) model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return model.Book{}
	}

	for _, book := range books {
		if book.ID == id {
			return book
		}
	}

	return model.Book{}
}

func (b *JSONBooks) GetBooksByAuthor(author string) map[int]model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}
	var formattedBooks = make(map[int]model.Book)
	for id, book := range books {
		if book.Author == author {
			formattedBooks[id] = book
		}
	}

	return formattedBooks
}

func (b *JSONBooks) UpdateBook(id int, title, author string) bool {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return false
	}

	for _, book := range books {
		if book.ID == id {
			books[id] = model.Book{ID: id, Title: title, Author: author}
			err = b.saveBooks(books)
			if err != nil {
				fmt.Printf("Failed to save books: %v\n", err)
				return false
			}
			return true
		}
	}

	return false
}

func (b *JSONBooks) DeleteBook(id int) bool {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return false
	}

	for i := range books {
		if i == id {
			delete(books, i)
			err = b.saveBooks(books)
			if err != nil {
				fmt.Printf("Failed to save books: %v\n", err)
				return false
			}
			return true
		}
	}

	return false
}
