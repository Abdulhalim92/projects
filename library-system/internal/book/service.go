package book

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	br BookRepository
}

func NewService(b BookRepository) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(b *model.Book) (*model.Book, error) {
	// books, err := s.br.GetBooksByAuthor(b.AuthorId)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(books) == 0 {
	// 	return nil, fmt.Errorf("Given author_id: %d doesn't exist", b.AuthorId)
	// }

	return s.br.AddBook(b)
}

func (s *Service) ListBooks() []model.Book {
	books, err := s.br.GetBooks()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}

func (s *Service) FindBook(id int) (*model.Book, bool) {
	book, err := s.br.GetBookByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return book, true
}

func (s *Service) FindBooksByAuthor(authorName string) []model.Book {
	author := model.Author{}
	err := s.br.db.First(&author, "name = ?", author).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	books, err := s.br.GetBooksByAuthor(author.AuthorID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}

func (s *Service) EditBook(id int, title, authorName string) bool {
	var author model.Author
	err := s.br.db.First(&author, "name = ?", authorName).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	book := model.Book{
		BookId:   id,
		Title:    title,
		AuthorId: author.AuthorID,
	}
	err = s.br.UpdateBook(&book)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) RemoveBook(id int) bool {
	err := s.br.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
