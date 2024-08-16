package user

import "projects/internal/model"

type UserInterface interface {
	AddUser(Password, UseerName string) model.User
	GetUsers() map[int]model.User
	GetUserById(id int) model.User
	DeleteUserById(id int) bool
}
