package book

import (
	"fmt"
	"log"
	"projects/internal/model"

	"gorm.io/gorm"
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
		log.Printf("Addbook: Failed to add book: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}

	return b, nil
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	var b []model.Book

	// select * from books;
	err := r.db.Find(&b).Error
	if err != nil {
		log.Printf("GetBooks: Failed to get books: %v\n", err)
		return nil, fmt.Errorf("Cannon find books with error: %v", err)
	}

	return b, nil
}

func (r *BookRepository) GetBookByID(bookID int) (*model.Book, error) {
	var b model.Book

	// select * from books where book_id = bookID
	err := r.db.Where("book_id = ?", bookID).Find(&b).Error
	if err != nil {
		log.Printf("GetBookByID: Failed to get book: %v\n", err)
		return nil, fmt.Errorf("Cannot find book with error: %v", err)
	}

	return &b, nil
}

func (r *BookRepository) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var b []model.Book

	// select * from books where author_id = authorID
	err := r.db.Where("author_id = ?", authorID).Find(&b).Error
	if err != nil {
		log.Printf("GetBooksByAuthor: Failed to get books: %v\n", err)
		return nil, fmt.Errorf("Cannot find books by authorID with error: %v", err)
	}

	return b, nil
}

func (r *BookRepository) UpdateBook(b *model.Book) (*model.Book, error) {
	// update books set title = 'War and Peace', author_id = 1 where book_id = 1
	result := r.db.Model(&b).Updates(&b)
	if result.Error != nil {
		log.Printf("UpdateBook: Failed to update book: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update book: %v\n", result.Error)
	}

	return b, nil
}

func (r *BookRepository) DeleteBook(bookID int) (int, error) {
	// delete from books where book_id = bookID returning book_id
	result := r.db.Delete(&model.Book{}, bookID)
	if result.Error != nil {
		log.Printf("DeleteBook: Failed to delete book: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete book: %v\n", result.Error)
	}

	return bookID, nil
}
