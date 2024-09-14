package service

import (
	"projects/internal/model"
	"projects/internal/utils"
)

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashPassword
	return s.Repository.AddUser(u)
}

func (s *Service) ListUsers() ([]*model.User, error) {
	return s.Repository.GetUsers()
}

func (s *Service) FindUser(id int) (*model.User, error) {
	return s.Repository.GetUserByID(id)
}

func (s *Service) EditUser(u *model.User) (*model.User, error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashPassword
	return s.Repository.UpdateUser(u)
}

func (s *Service) RemoveUser(id int) (int, error) {
	return s.Repository.DeleteUser(id)
}

func (s *Service) SignIn(user *model.User) (bool, error) {
	originalUser, err := s.FindUser(user.UserID)
	if err != nil {
		return false, err
	}
	if !utils.CheckPasswordHash(user.Password, originalUser.Password) {
		return false, nil
	}
	return true, nil
}
