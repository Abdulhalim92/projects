package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User
	// select * from users
	err := r.db.Find(&users).Error
	if err != nil {
		log.Printf("GetUsers: Error getting users: %v", err)
		return nil, err
	}
	if len(users) == 0 {
		log.Printf("GetUsers: No users found")
		return nil, fmt.Errorf("GetUsers: No roles found")
	}

	return users, nil
}

func (r *Repository) CreateUser(user *model.User) (int, error) {
	// insert into users (username, password, role_id) values ('admin', 'admin', 1) returning user_id
	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("CreateUser: Error creating user: %v", err)
		return 0, err
	}

	return user.UserID, nil
}
