package service

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/repository"
)

type BooksService struct {
	BookRepo repository.BooksRepo
}

func NewBookService(b repository.BooksRepo) *BooksService {
	return &BooksService{b}
}

func (s *BooksService) CreateBook(b *model.Book) (*model.Book, error) {

	books, err := s.BookRepo.GetBooksByAuthor(b.AuthorID)
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
	return s.BookRepo.AddBook(b)
}

func (s *BooksService) ListBooks() ([]model.Book, error) {
	return s.BookRepo.GetBooks()
}

func (s *BooksService) FindBook(bookID int) (*model.Book, error) {
	return s.BookRepo.GetBookByID(bookID)
}

func (s *BooksService) FindBooksByAuthor(authorID int) ([]model.Book, error) {
	return s.BookRepo.GetBooksByAuthor(authorID)
}

func (s *BooksService) EditBook(b *model.Book) (*model.Book, error) {
	return s.BookRepo.UpdateBook(b)
}

func (s *BooksService) RemoveBook(bookID int) (int, error) {
	return s.BookRepo.DeleteBook(bookID)
}
