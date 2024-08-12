package user

import "projects/internal/model"

type Service struct {
	Users Users
}

func NewService(b Users) *Service {
	return &Service{b}
}

func (s *Service) CreateUser(Login, Pasword string) model.User {
	return s.Users.AddUser(Login, Pasword)
}

func (s *Service) ListUser() []model.User {
	return s.Users.GetUser()
}

func (s *Service) FindUser(id int) *model.User {
	return s.Users.GetUserByID(id)
}

func (s *Service) FindUserByLogin(Login string) []model.User {
	return s.Users.GetUsersByLogin(Login)
}

func (s *Service) EditUser(id int, Login, Pasword string) bool {
	return s.Users.UpdateUser(id, Login, Pasword)
}

func (s *Service) RemoveUser(id int) bool {
	return s.Users.DeleteUser(id)
}
