package user

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

const userFile = "users.json"

type JSONUsers struct {
	filename string
}

func NewJSONUsers(filename string) *JSONUsers {
	return &JSONUsers{
		filename: filename,
	}
}

func (u *JSONUsers) loadUsers() (map[int]model.User, error) {
	var users = map[int]model.User{}
	err := utils.ReadJSONFromFile(u.filename, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *JSONUsers) saveUsers(users map[int]model.User) error {
	return utils.WriteJSONToFile(u.filename, &users)
}

func (u *JSONUsers) AddUser(username, password string) model.User {
	users, err := u.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load users: %v\n", err)
		return model.User{}
	}
	lastID := len(users)

	user := model.User{
		ID:       lastID,
		Username: username,
		Password: password,
	}

	users[lastID] = user

	err = u.saveUsers(users)
	if err != nil {
		fmt.Printf("Failed to save users: %v\n", err)
		return model.User{}
	}

	fmt.Printf("User with username %s and password %s is created\n", user.Username, user.Password)

	return user
}

func (u *JSONUsers) GetUsers() (map[int]model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %v\n", err)

	}

	return users, nil
}

func (u *JSONUsers) GetUserByID(id int) (*model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %w\n", err)
	}

	if user, ok := users[id]; ok {
		return &user, nil
	}

	return nil, fmt.Errorf("User with id: %d doesn't exist\n", id)
}

func (u *JSONUsers) GetUsersByUsername(username string) ([]model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %v\n", err)
	}

	var filteredUsers []model.User

	for _, user := range users {
		if user.Username == username {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers, nil
}

func (u *JSONUsers) UpdateUser(user model.User) error {
	users, err := u.loadUsers()
	if err != nil {
		return fmt.Errorf("Failed to load users: %v\n", err)
	}

	users[user.ID] = user
	err = u.saveUsers(users)
	if err != nil {
		return fmt.Errorf("Failed to save users: %v\n", err)
	}

	return nil
}

func (u *JSONUsers) DeleteUser(id int) error {
	users, err := u.loadUsers()
	if err != nil {
		return fmt.Errorf("Failed to load users: %v\n", err)
	}

	delete(users, id)
	err = u.saveUsers(users)
	if err != nil {
		return fmt.Errorf("Failed to save users: %v\n", err)

	}
	return nil
}
