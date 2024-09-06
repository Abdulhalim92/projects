package book

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	BookRepo Repository
}

func NewService(b Repository) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {
	books, err := s.BookRepo.GetBooksByAuthor(b.AuthorID)
	if err != nil {
		return nil, err
	}
	if len(books) > 0 {
		for _, book := range books {
			if book.Title == b.Title {
				return nil, fmt.Errorf("Books with authorID %d and title %s already exists.\n", b.AuthorID, b.Title)
			}
		}
	}
	return s.BookRepo.AddBook(b)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	return s.BookRepo.GetBooks()
}

func (s *Service) FindBook(bookID int) (*model.Book, error) {
	return s.BookRepo.GetBookByID(bookID)
}

func (s *Service) FindBooksByAuthor(authorID int) ([]model.Book, error) {
	return s.BookRepo.GetBooksByAuthor(authorID)
}

func (s *Service) EditBook(b *model.Book) (*model.Book, error) {
	return s.BookRepo.UpdateBook(b)
}

func (s *Service) RemoveBook(bookID int) (int, error) {
	return s.BookRepo.DeleteBook(bookID)
}
