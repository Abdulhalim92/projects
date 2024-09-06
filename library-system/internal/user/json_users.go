package user

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

type JSONUsers struct {
	fileName string
}

func NewJSONUsers(fileName string) *JSONUsers {
	return &JSONUsers{
		fileName: fileName,
	}
}

func (u *JSONUsers) loadUsers() (map[int]model.User, error) {
	var users map[int]model.User
	err := utils.ReadJSONFromFile(u.fileName, &users)
	return users, err
}

func (u *JSONUsers) saveUsers(users map[int]model.User) error {
	err := utils.WriteJSONToFile(u.fileName, &users)
	return err
}

func (u *JSONUsers) AddUser(username, password string) model.User {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load users: %v\n", err)
		return model.User{}
	}

	lastID := 0

	for id := range users {
		if id > lastID {
			lastID = id
		}
	}
	lastID++

	user := model.User{
		UserID:   lastID,
		Username: username,
		Password: password,
	}

	users[lastID] = user

	err = u.saveUsers(users)
	if err != nil {
		fmt.Printf("Failed to save users: %v\n", err)
		return model.User{}
	}

	fmt.Printf("User %s with password %s is created\n", user.Username, user.Password)

	return user
}

func (u *JSONUsers) GetUsers() map[int]model.User {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}
	return users
}

func (u *JSONUsers) GetUserByID(id int) *model.User {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load users: %v\n", err)
		return nil
	}

	user, exists := users[id]
	if !exists {
		return nil
	}
	return &user
}

func (u *JSONUsers) UpdateUser(id int, new_username, new_password string) bool {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load users: %v\n", err)
		return false
	}

	user, exists := users[id]
	if !exists {
		fmt.Printf("There is no user with id %d\n", id)
		return false
	}

	user.Username = new_username
	user.Password = new_password
	users[id] = user
	err = u.saveUsers(users)
	if err != nil {
		fmt.Printf("Failed to save users: %v\n", err)
		return false
	}

	return true
}

func (u *JSONUsers) DeleteUser(id int) bool {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load users: %v\n", err)
		return false
	}
	_, exists := users[id]

	if !exists {
		fmt.Printf("User with id %d not found\n", id)
		return false
	}

	delete(users, id)

	err = u.saveUsers(users)

	if err != nil {
		fmt.Printf("Failed to save users: %v\n", err)
		return false
	}

	return true
}
