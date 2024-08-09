package user

import (
	"fmt"
	"projects/internal/model"
)

type Users struct {
	UsersMap map[int]model.User
	lastID   int
}

func NewUsers(users map[int]model.User) *Users {
	return &Users{
		UsersMap: users,
		lastID:   0,
	}
}

var lastID int

func init() {
	lastID = 0
}

func (u *Users) AddUser(login, pswd string) model.User {
	lastID++
	user := model.User{
		ID:       lastID,
		Login:    login,
		Password: pswd,
	}

	u.UsersMap[lastID] = user

	fmt.Println("User with login %s and password %s is created", login, pswd)

	return user
}

func (u *Users) GetUsers() map[int]model.User {
	return u.UsersMap
}

func (u *Users) GetUsersByID(id int) *model.User {
	for k, user := range u.UsersMap {
		if k == id {
			fmt.Printf("Found user with id %d", id)
			return &user
		}

	}
	return nil
}

func (u *Users) GetUsersByLogin(login string) []model.User {
	var usersWithLogin []model.User

	for _, user := range u.UsersMap {
		if user.Login == login {
			usersWithLogin = append(usersWithLogin, user)
		}
	}

	return usersWithLogin
}

func (u *Users) UpdateUser(user model.User) bool {
	for k, _ := range u.UsersMap {
		if k == user.ID {
			u.UsersMap[k] = user
			return true
		}
	}
	return false
}

func (u *Users) DeleteUser(id int) bool {
	for k, _ := range u.UsersMap {
		if k == id {
			delete(u.UsersMap, id)
			return true
		}
	}
	return false
}
