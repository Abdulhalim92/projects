package service

import "projects/internal/model"

type ServiceInterface interface {
	CreateAuthor(a *model.Author) (*model.Author, error)
	GetAuthors() ([]model.Author, error)
	GetAuthorByID(authorID int) (*model.Author, error)
	EditAuthor(a *model.Author) (*model.Author, error)
	DeleteAuthor(authorID int) (int, error)
}
