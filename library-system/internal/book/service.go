package book

import "projects/internal/model"

type Service struct {
	Books Books
}

func NewService(b Books) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(title, author string) model.Book {
	return s.Books.AddBook(title, author)
}

func (s *Service) ListBooks() map[int]model.Book {
	return s.Books.GetBooks()
}

func (s *Service) FindBook(id int) *model.Book {
	return s.Books.GetBookByID(id)
}

func (s *Service) FindBooksByAuthor(author string) []model.Book {
	return s.Books.GetBooksByAuthor(author)
}

func (s *Service) EditBook(id int, title, author string) bool {
	book := model.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
	return s.Books.UpdateBook(book)
}

func (s *Service) RemoveBook(id int) bool {
	return s.Books.DeleteBook(id)
}
