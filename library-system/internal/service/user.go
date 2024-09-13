package service

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
	"projects/internal/utils"
)

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(u.UserID)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, err
	//}
	if err == gorm.ErrRecordNotFound {
		if err != nil {
			return nil, err
		}
	}
	if userByID != nil {
		return nil, fmt.Errorf("user with id %d already exists", u.UserID)
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		// TODO
	}

	u.Password = hashedPassword

	return s.Repository.AddUser(u)
}

func (s *Service) SignIn(u *model.User) (string, error) {
	user, err := s.Repository.GetUserByUsername(u.Username)
	if err == gorm.ErrRecordNotFound {
		if err != nil {
			return "", err
		}
	}

	b := utils.CheckPasswordHash(u.Password, user.Password)
	if !b {
		return "", fmt.Errorf("Введен неправильный пароль")
	}

	token, err := utils.GenerateJWT(u.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) ListUsers() ([]model.User, error) {
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return s.Repository.GetUsers()
}

func (s *Service) FindUser(id int) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return s.Repository.GetUserByID(id)
}

func (s *Service) EditUser(u *model.User) (*model.User, error) {
	userByID, err := s.Repository.GetUserByID(u.UserID)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", u.UserID)
	}

	return s.Repository.UpdateUser(u)
}

func (s *Service) DeleteUser(id int) (int, error) {
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	if userByID == nil {
		return 0, fmt.Errorf("user with id %d not found", id)
	}

	return s.Repository.DeleteUser(id)
}
