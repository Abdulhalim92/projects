package UserDataBase

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) AddUser(user *model.User) (*model.User, error) {
	err := u.Db.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error while adding a user: %v", err)
	}
	return user, nil
}
func (u *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := u.Db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting users: %v", err)
	}
	return users, nil
}
func (u *UserRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := u.Db.Table("users").Where("user_id = ?", id).Select("userID", "username", "password").Scan(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting a user: %v", err)
	}
	return &user, nil
}
func (u *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := u.Db.Updates(user).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a user: %v", err)
	}
	return user, nil
}
func (u *UserRepository) DeleteUserById(id int) (int, error) {
	err := u.Db.Table("users").Delete(&model.User{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error while deleting: %v", err)
	}
	return id, nil
}
