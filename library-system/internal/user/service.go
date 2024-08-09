package user

import "projects/internal/model"

type Service struct {
	Users Users
}

func NewService(u Users) *Service {
	return &Service{u}
}

func (s *Service) CreateUser(login, pswd string) model.User {
	return s.Users.AddUser(login, pswd)
}

func (s *Service) ListUsers() map[int]model.User {
	return s.Users.GetUsers()
}

func (s *Service) FindUser(id int) *model.User {
	return s.Users.GetUsersByID(id)
}

func (s *Service) FindUserByLogin(login string) []model.User {
	return s.Users.GetUsersByLogin(login)
}

func (s *Service) EditUser(id int, login, pswd string) bool {
	user := model.User{
		ID:       id,
		Login:    login,
		Password: pswd,
	}
	return s.Users.UpdateUser(user)
}

func (s *Service) RemoveUser(id int) bool {
	return s.Users.DeleteUser(id)
}
