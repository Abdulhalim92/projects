package service

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {

	authorByID, err := s.Repository.GetAuthorByID(b.AuthorID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if authorByID == nil {
		return nil, fmt.Errorf("Author with ID %d does not exist\n", b.AuthorID)
	}
	books, err := s.Repository.GetBooksByAuthor(b.AuthorID)
	if err != nil {
		return nil, err
	}
	if len(books) > 0 {
		for _, book := range books {
			if book.Title == b.Title {
				return nil, fmt.Errorf("BooksRepo with authorID %d and title %s already exists.\n", b.AuthorID, b.Title)
			}
		}
	}
	return s.Repository.AddBook(b)
}

func (s *Service) ListBooks() ([]*model.Book, error) {
	return s.Repository.GetBooks()
}

func (s *Service) FindBook(bookID int) (*model.Book, error) {
	return s.Repository.GetBookByID(bookID)
}

func (s *Service) FindBooksByAuthor(authorID int) ([]model.Book, error) {
	return s.Repository.GetBooksByAuthor(authorID)
}

func (s *Service) EditBook(b *model.Book) (*model.Book, error) {
	bookByID, err := s.Repository.GetBookByID(b.BookID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if bookByID == nil {
		return nil, fmt.Errorf("Book with ID %d does not exist\n", b.BookID)
	}
	return s.Repository.UpdateBook(b)
}

func (s *Service) RemoveBook(bookID int) (int, error) {
	return s.Repository.DeleteBook(bookID)
}
