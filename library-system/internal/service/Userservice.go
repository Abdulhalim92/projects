package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateUser(user *model.User) (*model.User, error) {
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		for _, userRow := range users {
			if user.Username == userRow.Username {
				return nil, fmt.Errorf("such username exists")
			}
		}
	}
	return s.Repository.AddUser(user)
}
func (s *Service) ListUsers() ([]model.User, error) {
	return s.Repository.GetUsers()
}
func (s *Service) ListUserById(id int) (*model.User, error) {
	user, err := s.Repository.GetUserById(id)
	if err != nil {
		return nil, err
	} else if user.UserID == 0 {
		return nil, fmt.Errorf("no user with such id")
	}
	return user, nil
}
func (s *Service) EditUser(NewUser *model.User) (*model.User, error) {
	user, err := s.Repository.GetUserById(NewUser.UserID)
	if err != nil {
		return nil, err
	} else if user.UserID == 0 {
		return nil, fmt.Errorf("no user with sich id")
	}
	return s.Repository.UpdateUser(NewUser)
}
func (s *Service) RemoveUser(id int) (int, error) {
	user, err := s.Repository.GetUserById(id)
	if err != nil {
		return 0, err
	} else if user.UserID == 0 {
		return 0, fmt.Errorf("no user with such id")
	}
	return s.Repository.DeleteUserById(id)
}
