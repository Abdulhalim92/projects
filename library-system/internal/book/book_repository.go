package book

import (
	"fmt"
	"library-system/internal/model"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
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
	var books []model.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookByID(bookID int) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, "id = ?", bookID).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var book []model.Book
	err := r.db.Find(&book, "author_id = ?", authorID).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(b *model.Book) error {
	err := r.db.Save(&model.Book{BookId: b.BookId, AuthorId: b.AuthorId, Title: b.Title}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) DeleteBook(bookID int) error {
	book, err := r.GetBookByID(bookID)
	if err != nil {
		return err
	}
	err = r.db.Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}
