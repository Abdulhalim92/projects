package user

import (
	"fmt"
	"projects/internal/model"
)

type Users struct {
	UsersMap map[int]model.User
	LastID   int
}

func NewUsers(users map[int]model.User) *Users {
	return &Users{
		UsersMap: users,
		LastID:   0,
	}
}

func (u *Users) AddUser(username, password string) model.User {
	u.LastID++
	user := model.User{
		ID:       u.LastID,
		Username: username,
		Password: password,
	}

	u.UsersMap[user.ID] = user

	fmt.Printf("User with username %s is created\n", user.Username)

	return user
}

func (u *Users) GetUsers() []model.User {
	var users []model.User
	for _, user := range u.UsersMap {
		users = append(users, user)
	}
	return users
}

func (u *Users) GetUserByID(id int) *model.User {
	user, exists := u.UsersMap[id]
	if !exists {
		fmt.Printf("User with id %d not found\n", id)
		return nil
	}
	return &user
}

func (u *Users) UpdateUser(id int, username, password string) bool {
	user, exists := u.UsersMap[id]
	if !exists {
		fmt.Printf("User with id %d not found\n", id)
		return false
	}

	user.Username = username
	user.Password = password

	u.UsersMap[id] = user

	fmt.Printf("User with id %d updated: Username: %s\n", id, user.Username)

	return true
}

func (u *Users) DeleteUser(id int) bool {
	_, exists := u.UsersMap[id]
	if !exists {
		fmt.Printf("User with id %d not found\n", id)
		return false
	}

	delete(u.UsersMap, id)

	return true
}
