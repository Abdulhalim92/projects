package book

import (
	"fmt"
	"projects/library-system/internal/model"
	"projects/library-system/internal/utils"
)

const booksFile = "books.json"

type JSONBooks struct {
	filename string
}

func NewJSONBooks(filename string) *JSONBooks {
	return &JSONBooks{
		filename: filename,
	}
}

func (b *JSONBooks) loadBooks() ([]model.Book, error) {
	var books []model.Book
	err := utils.ReadJSONFromFile(b.filename, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *JSONBooks) saveBooks(books []model.Book) error {
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
		if id > lastID {
			lastID = id
		}
	}
	lastID++

	book := model.Book{
		ID:     lastID,
		Title:  title,
		Author: author,
	}

	books = append(books, book)

	err = b.saveBooks(books)
	if err != nil {
		fmt.Printf("Failed to save books: %v\n", err)
		return model.Book{}
	}

	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)

	return book
}

func (b *JSONBooks) GetBooks() []model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	return books
}

func (b *JSONBooks) GetBookByID(id int) *model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}

	return nil
}

func (b *JSONBooks) GetBooksByAuthor(author string) []model.Book {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	var filteredBooks []model.Book

	for _, book := range books {
		if book.Author == author {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks
}

func (b *JSONBooks) UpdateBook(id int, title, author string) bool {
	books, err := b.loadBooks()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return false
	}

	for i, book := range books {
		if book.ID == id {
			books[i].Title = title
			books[i].Author = author
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

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
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