package service

import (
	"errors"
	"fmt"
	"projects/internal/model"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoBooksFound   = errors.New("no books found")
)

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {
	_, err := s.Repository.GetAuthorByID(b.AuthorID)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return nil, fmt.Errorf("author with id %d doesn't exists", b.AuthorID)
		}
		return nil, err
	}

	books, err := s.Repository.GetBooksByAuthor(b.AuthorID)
	if err != nil {
		return nil, err
	}

	//s.Repository.GetBookByAuthorAndID(b.AuthorID, b.BookID)

	if len(books) > 0 {
		for _, book := range books {
			if book.Title == b.Title {
				return nil, fmt.Errorf("the book whith title %s and authorID %d is already exists", b.Title, b.AuthorID)
			}
		}
	}

	return s.Repository.AddBook(b)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	books, err := s.Repository.GetBooks()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, ErrNoBooksFound
	}

	return s.Repository.GetBooks()
}

func (s *Service) FindBook(id int) (*model.Book, error) {
	book, err := s.Repository.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, fmt.Errorf("book with id %d not found", id)
	}

	return book, nil
}

func (s *Service) FindBooksByAuthor(authorID int) ([]model.Book, error) {
	bookByAuthor, err := s.Repository.GetBooksByAuthor(authorID)
	if err != nil {
		return nil, err
	}
	if len(bookByAuthor) == 0 {
		return nil, fmt.Errorf("no books with authorID %d", authorID)
	}

	return s.Repository.GetBooksByAuthor(authorID)
}

func (s *Service) EditBook(b *model.Book) (*model.Book, error) {
	_, err := s.Repository.GetBookByID(b.BookID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book with id %d not found", b.BookID)
		}
		return nil, err
	}

	return s.Repository.UpdateBook(b)
}

func (s *Service) RemoveBook(id int) (int, error) {
	_, err := s.Repository.GetBookByID(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("book with id %d not found", id)
		}
		return 0, err
	}

	return s.Repository.DeleteBook(id)
}

func (s *Service) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	booksByAuthor, err := s.Repository.GetBooksByAuthor(authorID)
	if err != nil {
		return nil, err
	}
	if len(booksByAuthor) == 0 {
		return nil, fmt.Errorf("no books with authorID %d", authorID)
	}

	return s.Repository.GetBooksByAuthor(authorID)
}
