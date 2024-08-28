package BookDataBase

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	b *BookRepository
}

func NewService(b *BookRepository) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	books, err := s.b.GetBooksByAuthor(book.Authorid)
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
	return s.b.AddBook(book)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	return s.b.GetBooks()
}
func (s *Service) FindBook(id int) (model.Book, error) {
	book, err := s.b.GetBookById(id)
	if err != nil {
		return model.Book{}, err
	} else if book == (model.Book{}) {
		return book, fmt.Errorf("book with such id doesn't exist")
	}
	return book, nil
}
func (s *Service) FindBookByAuthor(id int) ([]model.Book, error) {
	books, err := s.b.GetBooksByAuthor(id)
	if err != nil {
		return nil, err
	} else if len(books) == 0 {
		return nil, fmt.Errorf("no books with such author")
	}
	return books, nil
}
func (s *Service) EditBook(book model.Book) (bool, error) {
	bookById, err := s.b.GetBookById(book.Bookid)
	if err != nil {
		return false, err
	} else if bookById == (model.Book{}) {
		return false, fmt.Errorf("no booo with such id")
	}
	return s.b.UpdateBook(book)
}
func (s *Service) RemoveBook(id int) (bool, error) {
	book, err := s.b.GetBookById(id)
	if err != nil {
		return false, err
	} else if book == (model.Book{}) {
		return false, fmt.Errorf("no book with such id")
	}
	return s.b.DeleteBook(id)
}
