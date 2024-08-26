package book

import "projects/internal/model"

type Service struct {
	Books BookInterface
}

func NewService(b BookInterface) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(title string, author int) model.Book {
	return s.Books.AddBook(title, author)
}

func (s *Service) ListBooks() map[int]model.Book {
	return s.Books.GetBooks()
}

func (s *Service) FindBook(id int) model.Book {
	return s.Books.GetBookByID(id)
}

func (s *Service) FindBooksByAuthor(author int) map[int]model.Book {
	return s.Books.GetBooksByAuthor(author)
}

func (s *Service) EditBook(id int, title string, author int) bool {
	return s.Books.UpdateBook(id, title, author)
}

func (s *Service) RemoveBook(id int) bool {
	return s.Books.DeleteBook(id)
}
