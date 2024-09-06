package user

import (
	"projects/internal/model"
)

type Service struct {
	UsersRepo Repository
}

func NewService(u Repository) *Service {
	return &Service{u}
}

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	return s.UsersRepo.AddUser(u)
}

func (s *Service) ListUsers() ([]model.User, error) {
	return s.UsersRepo.GetUsers()
}

func (s *Service) FindUser(id int) (*model.User, error) {
	return s.UsersRepo.GetUserByID(id)
}

func (s *Service) EditUser(u *model.User) (*model.User, error) {
	return s.UsersRepo.UpdateUser(u)
}

func (s *Service) RemoveUser(id int) (int, error) {
	return s.UsersRepo.DeleteUser(id)
}
