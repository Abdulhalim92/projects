package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateBook(book *model.Book) (*model.Book, error) {
	books, err := s.Repository.GetBooksByAuthor(book.AuthorID)
	if err != nil {
		return nil, err
	}
	if len(books) > 0 {
		for _, bookAuthor := range books {
			if bookAuthor.Title == book.Title {
				return nil, fmt.Errorf("book already created")
			}
		}
	}
	return s.Repository.AddBook(book)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	return s.Repository.GetBooks()
}
func (s *Service) FindBook(id int) (*model.Book, error) {
	book, err := s.Repository.GetBookById(id)
	if err != nil {
		return nil, err
	}
	if book.BookID == 0 {
		return nil, fmt.Errorf("no such book")
	}
	return book, nil
}
func (s *Service) FindBooksByAuthor(id int) ([]model.Book, error) {
	books, err := s.Repository.GetBooksByAuthor(id)
	if err != nil {
		return nil, err
	} else if len(books) == 0 {
		return nil, fmt.Errorf("no books with such author id")
	}
	return books, nil
}
func (s *Service) EditBook(book *model.Book) (*model.Book, error) {
	bookById, err := s.Repository.GetBookById(book.BookID)
	if err != nil {
		return nil, err
	}
	if bookById.BookID == 0 {
		return nil, fmt.Errorf("no such book")
	}
	return s.Repository.UpdateBook(book)
}
func (s *Service) RemoveBook(id int) (int, error) {
	book, err := s.Repository.GetBookById(id)
	if err != nil {
		return 0, err
	}
	if book.BookID == 0 {
		return 0, fmt.Errorf("no book with such id")
	}
	return s.Repository.DeleteBook(id)
}
