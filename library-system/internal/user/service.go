package user

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	users JSONUsers
}

func NewService(u JSONUsers) *Service {
	return &Service{u}
}

func (s *Service) CreateUser(username, password string) (*model.User, bool) {
	user, err := s.users.AddUser(username, password)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return user, true
}

func (s *Service) ListUsers() []model.User {
	users, err := s.users.GetUsers()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}

func (s *Service) FindUser(id int) (*model.User, bool) {
	user, err := s.users.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return user, true
}

func (s *Service) FindUserByUsername(usrname string) (*model.User, bool) {
	user, err := s.users.GetUserByUsername(usrname)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return user, true
}

func (s *Service) EditUser(id int, username, password string) bool {
	user := model.User{
		ID:       id,
		Username: username,
		Password: password,
	}
	err := s.users.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) RemoveUser(id int) bool {
	err := s.users.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
