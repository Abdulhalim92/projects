package UserDataBase

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	U *UserRepository
}

func NewService(U *UserRepository) *Service {
	return &Service{U: U}
}

func (s *Service) CreateUser(user *model.User) (*model.User, error) {
	users, err := s.U.GetUsers()
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
	return s.U.AddUser(user)
}
func (s *Service) ListUsers() ([]model.User, error) {
	return s.U.GetUsers()
}
func (s *Service) ListUserById(id int) (*model.User, error) {
	user, err := s.U.GetUserById(id)
	if err != nil {
		return nil, err
	} else if user.UserID == 0 {
		return nil, fmt.Errorf("no user with such id")
	}
	return user, nil
}
func (s *Service) EditUser(NewUser *model.User) (*model.User, error) {
	user, err := s.U.GetUserById(NewUser.UserID)
	if err != nil {
		return nil, err
	} else if user.UserID == 0 {
		return nil, fmt.Errorf("no user with sich id")
	}
	return s.U.UpdateUser(NewUser)
}
func (s *Service) RemoveUser(id int) (int, error) {
	user, err := s.U.GetUserById(id)
	if err != nil {
		return 0, err
	} else if user.UserID == 0 {
		return 0, fmt.Errorf("no user with such id")
	}
	return s.U.DeleteUserById(id)
}
