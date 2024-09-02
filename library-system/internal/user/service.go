package user

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	ur UserRepo
}

func NewService(ur UserRepo) *Service {
	return &Service{ur}
}

func (s *Service) CreateUser(username, password string) (*model.User, error) {
	user := model.User{Username: username, Password: password}
	return s.ur.AddUser(&user)
}

func (s *Service) ListUsers() ([]model.User, error) {
	users, err := s.ur.GetUsers()
	if err != nil {

		return nil, fmt.Errorf("Error when listing the users: %e", err)
	}

	return users, nil
}

func (s *Service) FindUser(id int) (*model.User, error) {
	user, err := s.ur.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error occured when retrieiving user with id:%d\n%e", id, err)
	}

	return user, nil
}

func (s *Service) EditUser(id int, username, password string) error {
	user, err := s.FindUserByUsername(username)
	if err != nil {
		return err
	}

	err = s.ur.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("Error occured when editing user with id:%d\n%e", id, err)
	}
	return nil
}

func (s *Service) RemoveUser(id int) bool {
	err := s.ur.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) FindUserByUsername(name string) (*model.User, error) {
	var user model.User
	err := s.ur.db.First(&user, "username = ?", name).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't find user with username:%s\n%e", name, err)
	}

	return &user, nil
}
