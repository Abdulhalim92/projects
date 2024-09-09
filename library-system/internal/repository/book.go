package repository

import (
	"fmt"
	"projects/internal/model"
)

func (r *Repository) AddBook(book *model.Book) (*model.Book, error) {
	err := r.db.Create(book).Error
	if err != nil {
		return nil, fmt.Errorf("error adding the book: %v", err)
	}
	return book, nil
}

func (r *Repository) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, fmt.Errorf("error getting books: %v", err)
	}
	return books, nil
}

func (r *Repository) GetBookById(id int) (*model.Book, error) {
	var book model.Book
	err := r.db.Table("books").Where("book_id = ?", id).Scan(&book).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a book: %v", err)
	}
	return &book, nil
}

func (r *Repository) GetBooksByAuthor(AuthorId int) ([]model.Book, error) {
	var books []model.Book
	err := r.db.Where("books.authorID = ?", AuthorId).Find(&books).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting books by authorId %v", err)
	}
	return books, nil
}

func (r *Repository) UpdateBook(book *model.Book) (*model.Book, error) {
	err := r.db.Updates(book).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a book: %v", err)
	}
	return book, nil
}

func (r *Repository) DeleteBook(id int) (int, error) {
	err := r.db.Delete(model.Book{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error while deleting a book: %v", err)
	}
	return id, nil
}
