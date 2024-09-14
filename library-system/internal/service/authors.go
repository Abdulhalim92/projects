package service

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func (s *Service) ListAuthors() ([]*model.Author, error) {
	return s.Repository.GetAuthors()
}

func (s *Service) CreateAuthor(author *model.Author) (*model.Author, error) {
	authors, err := s.Repository.GetAuthors()
	if err != nil {
		return nil, err
	}
	for _, a := range authors {
		if a.Name == author.Name {
			return nil, fmt.Errorf("The author with this name %s already exists\n", author.Name)
		}
	}
	return s.Repository.AddAuthor(author)
}

func (s *Service) FindAuthorByID(id int) (*model.Author, error) {
	return s.Repository.GetAuthorByID(id)
}

func (s *Service) UpdateAuthor(author *model.Author) (*model.Author, error) {
	_, err := s.Repository.GetAuthorByID(author.AuthorID)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Errorf("Failed to Update Author - get author by id error %v\n", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		fmt.Errorf("Failed to Update Author - author with id %d does not exist\n", author.AuthorID)
		return nil, err
	}
	return s.Repository.EditAuthor(author)
}
