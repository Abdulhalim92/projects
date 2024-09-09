package repository

import (
	"fmt"
	"projects/internal/model"
)

func (r *Repository) AddUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error while adding a user: %v", err)
	}
	return user, nil
}
func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting users: %v", err)
	}
	return users, nil
}
func (r *Repository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Table("users").Where("user_id = ?", id).Select("userID", "username", "password").Scan(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting a user: %v", err)
	}
	return &user, nil
}
func (r *Repository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.db.Updates(user).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a user: %v", err)
	}
	return user, nil
}
func (r *Repository) DeleteUserById(id int) (int, error) {
	err := r.db.Table("users").Delete(&model.User{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error while deleting: %v", err)
	}
	return id, nil
}
