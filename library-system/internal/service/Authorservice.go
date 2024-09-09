package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateAuthor(author *model.Author) (*model.Author, error) {
	authors, err := s.Repository.GetAuthors()
	if err != nil {
		return nil, err
	}
	if len(authors) > 0 {
		for _, authorGot := range authors {
			if authorGot.Name == author.Name {
				return nil, fmt.Errorf("such author exists")
			}
		}
	}
	return s.Repository.AddAuthor(author)
}
func (s *Service) ListAuthors() ([]model.Author, error) {
	authors, err := s.Repository.GetAuthors()
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, fmt.Errorf("no authors exist")
	}
	return authors, nil
}
func (s *Service) ListAuthorById(id int) (*model.Author, error) {
	author, err := s.Repository.GetAuthorById(id)
	if err != nil {
		return nil, err
	} else if author.Authorid == 0 {
		return nil, fmt.Errorf("no author with such id")
	}
	return author, nil
}
func (s *Service) EditAuthor(a *model.Author) (*model.Author, error) {
	author, err := s.Repository.GetAuthorById(a.Authorid)
	if err != nil {
		return nil, err
	} else if author.Authorid == 0 {
		return nil, fmt.Errorf("no author with such id")
	}
	return s.Repository.UpdateAuthor(a)
}
func (s *Service) RemoveAuthor(id int) (int, error) {
	author, err := s.Repository.GetAuthorById(id)
	if err != nil {
		return 0, err
	} else if author.Authorid == 0 {
		return 0, fmt.Errorf("no author with such id")
	}
	return s.Repository.DeleteAuthor(id)
}
