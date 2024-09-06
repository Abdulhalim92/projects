package book

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddBook(b *model.Book) (*model.Book, error) {
	result := r.db.Create(&b)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}
	return b, nil
}

func (r *Repository) GetBooks() ([]model.Book, error) {
	var books []model.Book
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get books: %v\n", result.Error)
	}
	return books, nil
}

func (r *Repository) GetBookByID(bookID int) (*model.Book, error) {
	var book model.Book
	result := r.db.First(&book, bookID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get book by ID: %v\n", result.Error)
	}
	return &book, nil
}

func (r *Repository) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var books []model.Book
	result := r.db.Where("author_id = ?", authorID).Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get books by author ID: %v\n", result.Error)
	}
	return books, nil
}

func (r *Repository) UpdateBook(b *model.Book) (*model.Book, error) {
	var book model.Book
	result := r.db.First(&book, b.BookID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get book to update: %v\n", result.Error)
	}
	result = r.db.Save(b)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to update book: %v\n", result.Error)
	}
	return &book, nil
}

func (r *Repository) DeleteBook(bookID int) (int, error) {
	result := r.db.Delete(model.Book{BookID: bookID})
	if result.Error != nil {
		return 0, fmt.Errorf("Failed to delete book: %v\n", result.Error)
	}
	return bookID, nil
}
