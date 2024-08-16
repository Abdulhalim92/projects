package user

import "projects/internal/model"

type Service struct {
	Users UserInterface
}

func NewService(s UserInterface) *Service {
	return &Service{s}
}
func (s *Service) CreateUser(UserName string, password string) model.User {
	return s.Users.AddUser(UserName, password)
}
func (s Service) ListUsers() map[int]model.User {
	return s.Users.GetUsers()
}
func (s Service) ListUserById(id int) model.User {
	return s.Users.GetUserById(id)
}
func (s *Service) RemoveUserById(id int) bool {
	return s.Users.DeleteUserById(id)
}
