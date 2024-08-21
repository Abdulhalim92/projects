package book

import "projects/library-system/internal/model"


type Service struct {
	Books JSONBooks
}

func NewService(b JSONBooks) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(title, author string) model.Book {
	return s.Books.AddBook(title, author)
}

func (s *Service) ListBooks() []model.Book {
	return s.Books.GetBooks()
}

func (s *Service) FindBook(id int) *model.Book {
	return s.Books.GetBookByID(id)
}

func (s *Service) FindBooksByAuthor(author string) []model.Book {
	return s.Books.GetBooksByAuthor(author)
}

func (s *Service) EditBook(id int, title, author string) bool {
	return s.Books.UpdateBook(id, title, author)
}

func (s *Service) RemoveBook(id int) bool {
	return s.Books.DeleteBook(id)
}