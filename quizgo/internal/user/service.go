package user

import (
	"fmt"
	"quizgo/internal/model"
)

type Service struct {
	Users Users
}

func NewService(u *Users) *Service {
	return &Service{}
}

func (s *Service) CreateUser(login, pswd string) model.User {
	return s.Users.Add(login, pswd)
}

func (s *Service) ListUsers() map[int]model.User {
	return s.Users.GetUsers()
}

func (s *Service) FindUserByLogin(login string) (model.User, error) {
	if user, ok := s.Users.GetUserByLogin(login); !ok {
		return model.User{}, fmt.Errorf("no such login: %s", login)
	} else {
		return user, nil
	}
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
