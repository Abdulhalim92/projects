package book

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) AddBook(b *model.Book) (*model.Book, error) {
	result := r.db.Create(&b)
	// insert into books (title, author_id) values ('War and Peace', 1)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}

	return b, nil
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	// SELECT * FROM books;
	var books []model.Book

	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get books: %v\n", result.Error)
	}
	return books, nil
}

func (r *BookRepository) GetBookByID(bookID int) (*model.Book, error) {
	return nil, nil
}

func (r *BookRepository) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var books []model.Book

	result := r.db.Where("author_id = ?", authorID).Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get books: %v\n", result.Error)
	}
	return books, nil
}

func (r *BookRepository) UpdateBook(b *model.Book) (*model.Book, error) {
	return nil, nil
}

func (r *BookRepository) DeleteBook(bookID int) (int, error) {
	return 0, nil
}
