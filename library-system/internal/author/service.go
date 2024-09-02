package author

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	ar AuthorRepo
}

func NewService(ar AuthorRepo) *Service {
	return &Service{ar}
}

func (s *Service) CreateAuthor(name, bio string) (*model.Author, error) {
	author := model.Author{Name: name, Biography: bio}
	return s.ar.AddAuthor(&author)
}

func (s *Service) ListAuthors() ([]model.Author, error) {
	authors, err := s.ar.GetAuthors()
	if err != nil {

		return nil, fmt.Errorf("Error when listing the authors: %e", err)
	}

	return authors, nil
}

func (s *Service) FindAuthor(id int) (*model.Author, error) {
	author, err := s.ar.GetAuthorByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error occured when retrieiving author with id:%d\n%e", id, err)
	}

	return author, nil
}

func (s *Service) EditAuthor(id int, name, bio string) error {
	author, err := s.FindAuthorByName(name)
	if err != nil {
		return err
	}

	err = s.ar.UpdateAuthor(author)
	if err != nil {
		return fmt.Errorf("Error occured when editing author with id:%d\n%e", id, err)
	}
	return nil
}

func (s *Service) RemoveAuthor(id int) bool {
	err := s.ar.DeleteAuthor(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) FindAuthorByName(name string) (*model.Author, error) {
	var author model.Author
	err := s.ar.db.First(&author, "name = ?", name).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't find author with name:%s\n%e", name, err)
	}

	return &author, nil
}
