package user

import "projects/Medical_Records/internal/model"

type Service struct {
	Users Users
}

func CreateService(s Users) *Service {
	return &Service{s}
}
func (s *Service) CreateUser(id int, name string, surname string, age int) *model.User {
	return s.Users.AddUser(id, name, surname, age)
}
func (s Service) ListUsers() []model.User {
	return s.Users.GetUsers()
}
func (s Service) ListUserById(id int) *model.User {
	return s.Users.GetUserById(id)
}
func (s *Service) RemoveUserById(id int) bool {
	return s.Users.DeleteUserById(id)
}
