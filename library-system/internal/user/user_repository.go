package user

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddUser(u *model.User) (*model.User, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}
	return u, nil
}

func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*model.User, error) {
	var u model.User
	result := r.db.Find(&u, id)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get user by id: %v\n", result.Error)
	}
	return &u, nil
}

func (r *Repository) UpdateUser(u *model.User) (*model.User, error) {
	result := r.db.Save(u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}
	return u, nil
}

func (r *Repository) DeleteUser(id int) (int, error) {
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}
	return id, nil
}
