package book

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	BookRepository BookRepository
}

func NewService(b BookRepository) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {
	books, err := s.BookRepository.GetBooksByAuthor(b.AuthorID)
	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		for _, book := range books {
			if book.Title == b.Title {
				return nil, fmt.Errorf("the book whith title %s and authorID %d is already exists", b.Title, b.AuthorID)
			}
		}
	}

	return s.BookRepository.AddBook(b)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	return s.BookRepository.GetBooks()
}

func (s *Service) FindBook(id int) (*model.Book, error) {
	return s.BookRepository.GetBookByID(id)
}

func (s *Service) FindBooksByAuthor(authorID int) ([]model.Book, error) {
	return s.BookRepository.GetBooksByAuthor(authorID)
}

func (s *Service) EditBook(b *model.Book) (*model.Book, error) {
	book, err := s.BookRepository.GetBookByID(b.BookId)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, fmt.Errorf("book with id %d not found", b.BookId)
	}

	return s.BookRepository.UpdateBook(b)
}

func (s *Service) RemoveBook(id int) (int, error) {
	book, err := s.BookRepository.GetBookByID(id)
	if err != nil {
		return 0, err
	}
	if book == nil {
		return 0, fmt.Errorf("book with id %d not found", id)
	}

	return s.BookRepository.DeleteBook(id)
}
