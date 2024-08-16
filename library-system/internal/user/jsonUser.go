package user

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

const Filename = "user.json"

type JsonUser struct {
	filename string
}

func NewJSONUsers(filename string) *JsonUser {
	return &JsonUser{
		filename: filename,
	}
}
func (j JsonUser) loadUsers() (map[int]model.User, error) {
	users := make(map[int]model.User)
	err := utils.ReadJsonFromFileUser(j.filename, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (j JsonUser) saveUsers(users map[int]model.User) error {
	return utils.WriteJsonToUserFile(j.filename, &users)
}
func (j JsonUser) AddUser(Password, Login string) model.User {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error while loading Users: %v\n", err)
		return model.User{}
	}
	lastID := 0

	for id := range users {
		if id >= lastID {
			lastID = id
		}
	}
	lastID++
	users[lastID] = model.User{UserId: lastID, Password: Password, Login: Login}
	err = j.saveUsers(users)
	if err != nil {
		fmt.Printf("Error while saving: %v\n", err)
		return model.User{}
	}
	fmt.Printf("User with id: %d, Password: %s, Login: %s\n", lastID, Password, Login)
	return model.User{UserId: lastID, Password: Password, Login: Login}
}
func (j JsonUser) GetUsers() map[int]model.User {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error loading %v\n", err)
		return nil
	}
	return users
}
func (j JsonUser) GetUserById(id int) model.User {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error loading %v\n", err)
		return model.User{}
	}
	for id2 := range users {
		if id2 == id {
			return users[id]
		}
	}
	fmt.Printf("Book with such id is not found!!!\n")
	return model.User{}
}
func (j JsonUser) GetUsersByPassword(password string) map[int]model.User {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("error while loading users: %v", err)
		return nil
	}
	var formattedUsers = make(map[int]model.User)
	for _, user := range users {
		if user.Password == password {
			formattedUsers[user.UserId] = model.User{UserId: user.UserId, Password: user.Password, Login: user.Login}
		}
	}
	return formattedUsers
}
func (j JsonUser) UpdateUser(UserId int, Login, Password string) bool {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error while loading Users")
		return false
	}
	for id := range users {
		if users[id].UserId == UserId {
			users[id] = model.User{Login: Login, Password: Password}
			err = j.saveUsers(users)
			if err != nil {
				fmt.Printf("Error while saving to file: %v", err)
				return false
			}
			return true
		}
	}
	return false
}
func (j JsonUser) DeleteUser(id int) bool {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error while loading files: %v", err)
		return false
	}
	for i, user := range users {
		if user.UserId == id {
			delete(users, i)
			err = j.saveUsers(users)
			if err != nil {
				fmt.Printf("Error while saving to the file: %v", err)
				return false
			}
			return true
		}
	}
	return false
}
