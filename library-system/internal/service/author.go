package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateAuthor(a *model.Author) (*model.Author, error) {
	authorByName, err := s.Repository.GetAuthorByName(a.Name)
	if err != nil {
		return nil, err
	}

	if authorByName != nil {
		return nil, fmt.Errorf("author with name %s already exists", a.Name)
	}

	author, err := s.Repository.AddAuthor(a)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (s *Service) GetAuthors() ([]model.Author, error) {
	authors, err := s.Repository.GetAuthors()
	if err != nil {
		return nil, err
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("no authors found")
	}

	return authors, nil
}

func (s *Service) GetAuthorByID(authorID int) (*model.Author, error) {
	authorByID, err := s.Repository.GetAuthorByID(authorID)
	if err != nil {
		return nil, err
	}
	if authorByID == nil {
		return nil, fmt.Errorf("author with id %d not found", authorID)
	}

	return authorByID, nil
}

func (s *Service) EditAuthor(a *model.Author) (*model.Author, error) {
	authorByID, err := s.Repository.GetAuthorByID(a.AuthorID)
	if err != nil {
		return nil, err
	}
	if authorByID == nil {
		return nil, fmt.Errorf("author with id %d not found", a.AuthorID)
	}

	return s.Repository.UpdateAuthor(a)
}

func (s *Service) DeleteAuthor(authorID int) (int, error) {
	authorByID, err := s.Repository.GetAuthorByID(authorID)
	if err != nil {
		return 0, err
	}
	if authorByID == nil {
		return 0, fmt.Errorf("author with id %d not found", authorID)
	}

	return s.Repository.DeleteAuthor(authorID)
}
