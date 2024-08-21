package user

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
	"sort"
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

func (u *JSONUsers) AddUser(username, password string) (*model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %e\n", err)
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
		return nil, fmt.Errorf("Failed to save users: %e\n", err)
	}

	fmt.Printf("User with username %s and password %s is created\n", user.Username, user.Password)

	return &user, nil
}

func (u *JSONUsers) GetUsers() ([]model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %v\n", err)
	}

	var sliceUsers []model.User

	for _, v := range users {
		sliceUsers = append(sliceUsers, v)
	}

	sort.Slice(sliceUsers, func(i, j int) bool {
		if sliceUsers[i].ID < sliceUsers[j].ID {
			return true
		}
		return false
	})

	return sliceUsers, nil
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

func (u *JSONUsers) GetUserByUsername(usrname string) (*model.User, error) {
	users, err := u.loadUsers()
	if err != nil {
		return nil, fmt.Errorf("Failed to load users: %e\n", err)
	}

	for _, user := range users {
		if user.Username == usrname {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("Couldn't find user with username:%v", usrname)
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
