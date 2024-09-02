package BookDataBase

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (b *BookRepository) AddBook(book *model.Book) (*model.Book, error) {
	err := b.db.Raw("INSERT INTO books (title, author_id)VALUES ('?', '?')", book.Title, book.AuthorID).Error
	if err != nil {
		return nil, fmt.Errorf("error adding the book: %v", err)
	}
	return book, nil
}

func (b *BookRepository) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := b.db.Find(&books).Error
	if err != nil {
		return nil, fmt.Errorf("error getting books: %v", err)
	}
	return books, nil
}

func (b *BookRepository) GetBookById(id int) (*model.Book, error) {
	var book model.Book
	err := b.db.Table("books").Where("bookID = ?", id).Scan(&book).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a book: %v", err)
	}
	return &book, nil
}

func (b *BookRepository) GetBooksByAuthor(AuthorId int) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Where("books.authorID = ?", AuthorId).Find(&books).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting books by authorId %v", err)
	}
	return books, nil
}

func (b *BookRepository) UpdateBook(book *model.Book) (*model.Book, error) {
	err := b.db.Updates(book).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a book: %v", err)
	}
	return book, nil
}

func (b *BookRepository) DeleteBook(id int) (int, error) {
	err := b.db.Delete(&model.Book{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error while deleting a book: %v", err)
	}
	return id, nil
}
