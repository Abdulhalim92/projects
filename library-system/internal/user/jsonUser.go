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
	err := utils.ReadJSONFromFile(j.filename, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (j JsonUser) saveUsers(users map[int]model.User) error {
	return utils.WriteJSONToFile(j.filename, &users)
}
func (j JsonUser) AddUser(Password, UserName string) model.User {
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
	users[lastID] = model.User{Users_id: lastID, Password: Password, Username: UserName}
	err = j.saveUsers(users)
	if err != nil {
		fmt.Printf("Error while saving: %v\n", err)
		return model.User{}
	}
	fmt.Printf("User with id: %d, Password: %s, Login: %s\n", lastID, Password, UserName)
	return model.User{Users_id: lastID, Password: Password, Username: UserName}
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
			formattedUsers[user.Users_id] = model.User{Users_id: user.Users_id, Password: user.Password, Username: user.Username}
		}
	}
	return formattedUsers
}
func (j JsonUser) UpdateUser(UserId int, UserName, Password string) bool {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error while loading Users")
		return false
	}
	for id := range users {
		if users[id].Users_id == UserId {
			users[id] = model.User{Username: UserName, Password: Password}
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
func (j JsonUser) DeleteUserById(id int) bool {
	users, err := j.loadUsers()
	if err != nil {
		fmt.Printf("Error while loading files: %v", err)
		return false
	}
	for i, user := range users {
		if user.Users_id == id {
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
