package service

import (
	"projects/internal/repository"
)

type Service struct {
	Repository repository.Repository
}

func NewService(b repository.Repository) *Service {
	return &Service{b}
}
