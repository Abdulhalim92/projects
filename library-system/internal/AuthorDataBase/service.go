package AuthorDataBase

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	a *AuthorRep
}

func CreateService(a *AuthorRep) *Service {
	return &Service{a: a}
}
func (s *Service) CreateAuthor(author *model.Author) (*model.Author, error) {
	authors, err := s.a.GetAuthors()
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
	return s.a.AddAuthor(author)
}
func (s *Service) ListAuthors() ([]model.Author, error) {
	authors, err := s.a.GetAuthors()
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, fmt.Errorf("no authors exist")
	}
	return authors, nil
}
func (s *Service) ListAuthorById(id int) (*model.Author, error) {
	author, err := s.a.GetAuthorById(id)
	if err != nil {
		return nil, err
	} else if author.Authorid == 0 {
		return nil, fmt.Errorf("no author with such id")
	}
	return author, nil
}
func (s *Service) EditAuthor(a *model.Author) (*model.Author, error) {
	author, err := s.a.GetAuthorById(a.Authorid)
	if err != nil {
		return nil, err
	} else if author.Authorid == 0 {
		return nil, fmt.Errorf("no author with such id")
	}
	return s.a.UpdateAuthor(a)
}
func (s *Service) RemoveAuthor(id int) (int, error) {
	author, err := s.a.GetAuthorById(id)
	if err != nil {
		return 0, err
	} else if author.Authorid == 0 {
		return 0, fmt.Errorf("no author with such id")
	}
	return s.a.DeleteAuthor(id)
}
