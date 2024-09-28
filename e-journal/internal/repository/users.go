package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
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
func (r *Repository) GetUserByID(userID int) (*model.User, error) {
	var user *model.User
	// select * from users where user_id = ?
	err := r.db.First(&user, userID).Error
	if err != nil {
		log.Printf("GetUserByID: Error getting user by ID: %v", err)
		return nil, err
	}
	return user, nil
}
func (r *Repository) UpdateUser(user *model.User) (*model.User, error) {
	// update users set name = 'admin' where user_id = 1
	err := r.db.Clauses(clause.Returning{}).Updates(user).Error
	if err != nil {
		log.Printf("UpdateUser: Error updating user: %v", err)
		return nil, err
	}

	return user, nil
}
func (r *Repository) DeleteUser(userID int) error {
	// delete from users where user_id = 1
	err := r.db.Delete(&model.User{}, userID).Error
	if err != nil {
		log.Printf("DeleteUser: Error deleting user: %v", err)
		return err
	}

	return nil
}
