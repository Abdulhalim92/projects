package service

import (
	"projects/internal/model"
	"projects/internal/repository"
	"projects/internal/utils"
)

type UserService struct {
	UsersRepo repository.UsersRepo
}

func NewUserService(u repository.UsersRepo) *UserService {
	return &UserService{u}
}

func (s *UserService) CreateUser(u *model.User) (*model.User, error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashPassword
	return s.UsersRepo.AddUser(u)
}



func (s *UserService) ListUsers() ([]model.User, error) {
	return s.UsersRepo.GetUsers()
}

func (s *UserService) FindUser(id int) (*model.User, error) {
	return s.UsersRepo.GetUserByID(id)
}

func (s *UserService) EditUser(u *model.User) (*model.User, error) {
	return s.UsersRepo.UpdateUser(u)
}

func (s *UserService) RemoveUser(id int) (int, error) {
	return s.UsersRepo.DeleteUser(id)
}
