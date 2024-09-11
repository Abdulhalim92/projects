package repository

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) AddUser(u *model.User) (*model.User, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}
	return u, nil
}

func (r *UsersRepo) GetUsers() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *UsersRepo) GetUserByID(id int) (*model.User, error) {
	var u model.User
	result := r.db.Find(&u, id)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get user by id: %v\n", result.Error)
	}
	return &u, nil
}

func (r *UsersRepo) UpdateUser(u *model.User) (*model.User, error) {
	result := r.db.Save(u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}
	return u, nil
}

func (r *UsersRepo) DeleteUser(id int) (int, error) {
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}
	return id, nil
}
