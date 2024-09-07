package author

import (
	"projects/internal/model"
)

type Service interface {
	AddAuthor(a *model.Author) (*model.Author, error)
	GetAuthors() ([]model.Author, error)
	GetAuthorByID(authorID int) (*model.Author, error)
	UpdateAuthor(a *model.Author) (*model.Author, error)
	DeleteAuthor(authorID int) (int, error)
}

type ServiceImpl struct {
	AuthorRepository AuthorRepository
}

func NewAuthorService(ar AuthorRepository) *ServiceImpl {
	return &ServiceImpl{
		AuthorRepository: ar,
	}
}

func (s *ServiceImpl) AddAuthor(a *model.Author) (*model.Author, error) {
	return s.AuthorRepository.AddAuthor(a)
}

func (s *ServiceImpl) GetAuthors() ([]model.Author, error) {
	return s.AuthorRepository.GetAuthors()
}

func (s *ServiceImpl) GetAuthorByID(authorID int) (*model.Author, error) {
	return s.AuthorRepository.GetAuthorByID(authorID)
}

func (s *ServiceImpl) UpdateAuthor(a *model.Author) (*model.Author, error) {
	return s.AuthorRepository.UpdateAuthor(a)
}

func (s *ServiceImpl) DeleteAuthor(authorID int) (int, error) {
	return s.AuthorRepository.DeleteAuthor(authorID)
}
