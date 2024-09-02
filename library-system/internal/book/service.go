package book

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	br BookRepo
}

func NewService(b BookRepo) *Service {
	return &Service{b}
}

func (s *Service) CreateBook(title string, authorId int) (*model.Book, error) {
	// books, err := s.br.GetBooksByAuthor(b.AuthorId)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(books) == 0 {
	// 	return nil, fmt.Errorf("Given author_id: %d doesn't exist", b.AuthorId)
	// }
	book := model.Book{Title: title, AuthorId: authorId}

	return s.br.AddBook(&book)
}

func (s *Service) ListBooks() ([]model.Book, error) {
	books, err := s.br.GetBooks()
	if err != nil {
		return nil, fmt.Errorf("Error when listing the books: %e", err)
	}

	return books, nil
}

func (s *Service) FindBook(id int) (*model.Book, error) {
	book, err := s.br.GetBookByID(id)
	if err != nil {
		return nil, fmt.Errorf("Error occured when retrieiving book with id:%d\n%e", id, err)
	}

	return book, nil
}

func (s *Service) FindBooksByAuthor(authorName string) ([]model.Book, error) {
	author, err := s.FindAuthorByName(authorName)
	if err != nil {
		return nil, err
	}

	books, err := s.br.GetBooksByAuthor(author.AuthorId)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't get books of author with id:%d\n%e", author.AuthorId, err)
	}

	return books, nil
}

func (s *Service) EditBook(id int, title, authorName string) error {
	author, err := s.FindAuthorByName(authorName)
	if err != nil {
		return err
	}
	book := model.Book{
		BookId:   id,
		Title:    title,
		AuthorId: author.AuthorId,
	}
	err = s.br.UpdateBook(&book)
	if err != nil {
		return fmt.Errorf("Error occured when editing book with id:%d\n%e", id, err)
	}
	return nil
}

func (s *Service) RemoveBook(id int) bool {
	err := s.br.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) FindAuthorByName(name string) (*model.Author, error) {
	author := model.Author{}
	err := s.br.db.First(&author, "name = ?", name).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't find author with name:%s\n%e", name, err)
	}

	return &author, nil
}
