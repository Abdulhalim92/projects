package user

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	ur UserRepository
}

func NewService(ur UserRepository) *Service {
	return &Service{ur}
}

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	return s.ur.AddUser(u)
}

func (s *Service) ListUsers() []model.User {
	users, err := s.ur.GetUsers()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return users
}

func (s *Service) FindUser(id int) (*model.User, bool) {
	user, err := s.ur.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return user, true
}

func (s *Service) EditUser(id int, username, password string) bool {
	user, err := s.ur.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = s.ur.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) RemoveUser(id int) bool {
	err := s.ur.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
