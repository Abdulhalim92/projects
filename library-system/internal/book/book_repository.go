package book

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) AddBook(b *model.Book) (*model.Book, error) {
	result := r.db.Create(&b)
	// insert into books (title, author_id) values ('War and Peace', 1)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}

	return b, nil
}

func (r *BookRepo) GetBooks() ([]model.Book, error) {
	var books []model.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) GetBookByID(bookID int) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, "id = ?", bookID).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepo) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var book []model.Book
	err := r.db.Find(&book, "author_id = ?", authorID).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepo) UpdateBook(b *model.Book) error {
	result := r.db.Model(&b).Updates(&b)
	if result.Error != nil {
		log.Printf("UpdateBook: Failed to update book: %v\n", result.Error)
		return fmt.Errorf("Failed to update book: %v\n", result.Error)
	}

	return nil
}

func (r *BookRepo) DeleteBook(bookID int) error {
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
