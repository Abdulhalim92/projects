package user

import (
	"fmt"
	"projects/internal/model"
	"projects/internal/utils"
)

const booksFile = "user.json"

type JSONUsers struct {
	filename string
}

func NewJSONUsers(filename string) *JSONUsers {
	return &JSONUsers{
		filename: filename,
	}
}

func (b *JSONUsers) loadUsers() ([]model.User, error) {
	var users []model.User
	err := utils.ReadJSONFromFile(b.filename, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (b *JSONUsers) saveUsers(users []model.User) error {
	return utils.WriteJSONToFile(b.filename, &users)
}

func (b *JSONUsers) AddUser(Login, Pasword string) model.User {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
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
		ID:      lastID,
		Login:   Login,
		Pasword: Pasword,
	}

	users = append(users, user)

	err = b.saveUsers(users)
	if err != nil {
		fmt.Printf("Failed to save books: %v\n", err)
		return model.User{}
	}

	fmt.Printf("Book with tittle %s and author %s is created\n", user.Login, user.Pasword)

	return user
}

func (b *JSONUsers) GetUsers() []model.User {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	return users
}

func (b *JSONUsers) GetUserByID(id int) *model.User {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func (b *JSONUsers) GetUsersByLogin(Login string) []model.User {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return nil
	}

	var filteredUsers []model.User

	for _, user := range users {
		if user.Login == Login {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers
}

func (b *JSONUsers) UpdateUser(id int, Login, Pasword string) bool {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return false
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Login = Login
			users[i].Pasword = Pasword
			err = b.saveUsers(users)
			if err != nil {
				fmt.Printf("Failed to save books: %v\n", err)
				return false
			}
			return true
		}
	}

	return false
}

func (b *JSONUsers) DeleteUser(id int) bool {
	users, err := b.loadUsers()
	if err != nil {
		fmt.Printf("Failed to load books: %v\n", err)
		return false
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			err = b.saveUsers(users)
			if err != nil {
				fmt.Printf("Failed to save books: %v\n", err)
				return false
			}
			return true
		}
	}

	return false
}
