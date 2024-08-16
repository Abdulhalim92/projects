package book

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	Books JSONBooks
}

func NewService(b JSONBooks) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(title, author string) model.Book {
	return s.Books.AddBook(title, author)
}

func (s *Service) ListBooks() map[int]model.Book {
	books, err := s.Books.GetBooks()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return books
}

func (s *Service) FindBook(id int) *model.Book {
	book, err := s.Books.GetBookByID(id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return book
}

func (s *Service) FindBooksByAuthor(author string) []model.Book {
	books, err := s.Books.GetBooksByAuthor(author)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return books
}

func (s *Service) EditBook(id int, title, author string) bool {
	_, err := s.Books.GetBookByID(id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	book := model.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	err = s.Books.UpdateBook(book)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (s *Service) RemoveBook(id int) bool {
	err := s.Books.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
