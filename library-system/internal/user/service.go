package user

import (
	"projects/internal/model"
)

type Service struct {
	UserRepository UserRepository
}

func NewService(u UserRepository) *Service {
	return &Service{
		UserRepository: u,
	}
}

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	return s.UserRepository.AddUser(u)
}

func (s *Service) ListUsers() ([]model.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *Service) FindUser(id int) (*model.User, error) {
	return s.UserRepository.GetUserByID(id)
}

func (s *Service) EditUser(u *model.User) (*model.User, error) {
	return s.UserRepository.UpdateUser(u)
}

func (s *Service) RemoveUser(id int) (int, error) {
	return s.UserRepository.DeleteUser(id)
}
