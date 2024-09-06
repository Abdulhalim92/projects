package user

import (
	"fmt"

	"projects/internal/model"
)

type Users struct {
	UsersMap map[int]model.User
	Lastid   int
}

func NewUsers(users map[int]model.User) *Users {
	return &Users{
		UsersMap: users,
		Lastid:   0,
	}
}

func (b *Users) AddUser(Login, Pasword string) model.User {
	b.Lastid++
	user := model.User{
		ID:      b.Lastid,
		Login:   Login,
		Pasword: Pasword,
	}
	b.UsersMap[b.Lastid] = user

	fmt.Printf("User with Login %s and Pasword %s is created\n", user.Login, user.Pasword)

	return b.UsersMap[b.Lastid]
}

func (b *Users) GetUser() []model.User {
	users := make([]model.User, 0)
	for _, value := range b.UsersMap {
		users = append(users, value)
	}
	return users
}

func (b *Users) GetUserByID(id int) *model.User {
	value, ok := b.UsersMap[id]
	if !ok {
		fmt.Printf("Does't exist\n")
		return nil
	}
	return &value
}

func (b *Users) GetUsersByLogin(Login string) []model.User {
	var usersByLogin []model.User
	for _, value := range b.UsersMap {
		if value.Login == Login {
			usersByLogin = append(usersByLogin, value)
		}
	}
	return usersByLogin
}

func (b *Users) UpdateUser(id int, Login, Pasword string) bool {
	for key := range b.UsersMap {
		if key == id {
			b.UsersMap[key] = model.User{Login: Login, Pasword: Pasword}
			return true
		}
	}
	return false
}

func (b *Users) DeleteUser(id int) bool {
	for key := range b.UsersMap {
		if key == id {
			delete(b.UsersMap, key)
			return true
		}
	}
	return false
}
