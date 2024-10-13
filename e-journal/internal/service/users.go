package service

import (
	"errors"
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

var (
	ErrNotFound = errors.New("record not found")
)

func (s *Service) CreateUser(user *model.User) (int, error) {
	_, err := s.Repository.GetUserByName(user.Username)
	if err != nil && !errors.As(err, &ErrNotFound) {
		return 0, err
	} else if errors.As(err, &ErrNotFound) {
		return 0, fmt.Errorf("user with username %s is already exists", user.Username)
	}

	createdUserID, err := s.Repository.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return createdUserID, nil
	// Хешируем пароль пользователя
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Присваиваем хэшированный пароль пользователю
	user.Password = hashedPassword

	// Добавляем пользователя в репозиторий
	return s.Repository.CreateUser(user)
}

func (s *Service) UpdateUser(user *model.User) (*model.User, error) {
	_, err := s.Repository.GetUserByName(user.Username)
	if err != nil {
		if errors.As(err, &ErrNotFound) {
			return nil, fmt.Errorf("user with username %s is not exist", user.Username)
		}
		return nil, err
	}

	updatedUser, err := s.Repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) Login(user *model.User) (string, error) {

}
