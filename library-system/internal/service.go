package internal

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

// BookService

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {
	books, err := s.Repository.GetBooksByAuthor(b.AuthorID)
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

	return s.Repository.AddBook(b)
}

func (s *Service) ListBooks() ([]model.Book, error) {
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
	book, err := s.Repository.GetBookByID(b.BookId)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, fmt.Errorf("book with id %d not found", b.BookId)
	}

	return s.Repository.UpdateBook(b)
}

func (s *Service) RemoveBook(id int) (int, error) {
	book, err := s.Repository.GetBookByID(id)
	if err != nil {
		return 0, err
	}
	if book == nil {
		return 0, fmt.Errorf("book with id %d not found", id)
	}

	return s.Repository.DeleteBook(id)
}

func (s *Service) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	return s.Repository.GetBooksByAuthor(authorID)
}

// UserService

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(u.UserID)
	if err != nil {
		return nil, err
	}
	if userByID != nil {
		return nil, fmt.Errorf("user with id %d already exists", u.UserID)
	}

	return s.Repository.AddUser(u)
}

func (s *Service) ListUsers() ([]model.User, error) {
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return s.Repository.GetUsers()
}

func (s *Service) FindUser(id int) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return s.Repository.GetUserByID(id)
}

func (s *Service) EditUser(u *model.User) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(u.UserID)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", u.UserID)
	}

	return s.Repository.UpdateUser(u)
}

func (s *Service) RemoveUser(id int) (int, error) {
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	if userByID == nil {
		return 0, fmt.Errorf("user with id %d not found", id)
	}

	return s.Repository.DeleteUser(id)
}
