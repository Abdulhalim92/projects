package service

import (
	"errors"
	"fmt"
	"log"
	"projects/internal/model"
	"projects/internal/utils"
)

// CreateUser функция, которая добавляет пользователя в репозиторий.
func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	err := u.ValidateUser()
	if err != nil {
		return nil, err
	}
	// Проверка существования пользователя по ID
	userByID, err := s.Repository.GetUserByUsername(u.Username)

	// Если ошибка не nil и это не ошибка "record not found", возвращаем её
	if err != nil && !errors.As(err, &ErrRecordNotFound) {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	// Если пользователь существует (userByID не nil), возвращаем ошибку
	if userByID != nil {
		log.Printf("user with name %s is already exists", u.Username)
		return nil, fmt.Errorf("user with name %s is already exists", u.Username)
	}

	// Хешируем пароль пользователя
	hashedPassword, err := utils.HashPassword(*u.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Присваиваем хэшированный пароль пользователю
	u.Password = &hashedPassword

	// Добавляем пользователя в репозиторий
	return s.Repository.AddUser(u)
}

// SignIn функция, которая возвращает JWT токен для авторизованного пользователя.
func (s *Service) SignIn(u *model.User) (string, error) {
	// Проверка существования пользователя
	user, err := s.Repository.GetUserByUsername(u.Username)

	// Если ошибка "record not found", то пользователь не существует
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return "", fmt.Errorf("пользователь с именем %s не найден", u.Username)
		}
		// Если другая ошибка, вернуть её
		return "", err
	}

	// Проверка пароля пользователя
	if !utils.CheckPasswordHash(*u.Password, *user.Password) {
		return "", fmt.Errorf("введен неправильный пароль")
	}

	// Генерация JWT токена для авторизованного пользователя
	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", fmt.Errorf("не удалось сгенерировать токен: %w", err)
	}

	return token, nil
}

// ListUsers функция, которая возвращает список всех пользователей из репозитория.
func (s *Service) ListUsers() ([]model.User, error) {
	// Получаем список пользователей из репозитория
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}

	// Проверяем, что список пользователей не пустой
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return users, nil
}

// FindUser функция, которая возвращает пользователя по ID.
func (s *Service) FindUser(id int) (*model.User, error) {
	// Получаем пользователя по ID из репозитория
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors.As(err, &ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return userByID, nil
}

// EditUser функция, которая редактирует пользователя в репозитории.
func (s *Service) EditUser(u *model.User) (*model.User, error) {
	// Получаем пользователя по ID из репозитория
	_, err := s.Repository.GetUserByID(u.UserID)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", u.UserID)
		}
		return nil, err
	}

	// Обновляем информацию о пользователе
	updatedUser, err := s.Repository.UpdateUser(u)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Не отправляем пароль в ответ
	if *updatedUser.Password != "" {
		updatedUser.Password = nil
	}

	return updatedUser, nil
}

// DeleteUser функция, которая удаляет пользователя из репозитория.
func (s *Service) DeleteUser(id int) (int, error) {
	// Получаем пользователя по ID из репозитория
	_, err := s.Repository.GetUserByID(id)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("user with id %d not found", id)
		}
		return 0, err
	}

	// Удаляем пользователя
	deletedRows, err := s.Repository.DeleteUser(id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user with id %d: %w", id, err)
	}

	return deletedRows, nil
}
