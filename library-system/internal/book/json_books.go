package book

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
	"sort"
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

func (b *JSONBooks) loadBooks() (map[int]model.Book, error) {
	var books = map[int]model.Book{}
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
	lastID := len(books)

	// for id := range books {
	// 	if id > lastID {
	// 		lastID = id
	// 	}
	// }
	// lastID++

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

func (b *JSONBooks) GetBooks() ([]model.Book, error) {
	books, err := b.loadBooks()

	if err != nil {
		return nil, fmt.Errorf("Failed to load books: %v\n", err)

	}
	var sliceBooks []model.Book
	for _, v := range books {
		sliceBooks = append(sliceBooks, v)
	}

	sort.Slice(sliceBooks, func(i, j int) bool {
		if sliceBooks[i].ID < sliceBooks[j].ID {
			return true
		}
		return false
	})

	return sliceBooks, nil
}

func (b *JSONBooks) GetBookByID(id int) (*model.Book, error) {
	books, err := b.loadBooks()
	if err != nil {
		return nil, fmt.Errorf("Failed to load books: %w\n", err)
	}

	if book, ok := books[id]; ok {
		return &book, nil
	}

	return nil, fmt.Errorf("Book with id: %d doesn't exist\n", id)
}

func (b *JSONBooks) GetBooksByAuthor(author string) ([]model.Book, error) {
	books, err := b.loadBooks()
	if err != nil {
		return nil, fmt.Errorf("Failed to load books: %v\n", err)
	}

	var filteredBooks []model.Book

	for _, book := range books {
		if book.Author == author {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks, nil
}

func (b *JSONBooks) UpdateBook(book model.Book) error {
	books, err := b.loadBooks()
	if err != nil {
		return fmt.Errorf("Failed to load books: %v\n", err)
	}

	books[book.ID] = book
	err = b.saveBooks(books)
	if err != nil {
		return fmt.Errorf("Failed to save books: %v\n", err)
	}

	return nil
}

func (b *JSONBooks) DeleteBook(id int) error {
	books, err := b.loadBooks()
	if err != nil {
		return fmt.Errorf("Failed to load books: %v\n", err)
	}

	delete(books, id)
	err = b.saveBooks(books)
	if err != nil {
		return fmt.Errorf("Failed to save books: %v\n", err)

	}
	return nil
}
