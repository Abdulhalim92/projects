package user

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) AddUser(u *model.User) (*model.User, error) {
	// insert into users (username, password) values ('admin', 'admin')
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Printf("AddUser: Failed to add user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}

	return u, nil
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User

	// select * from users
	result := r.db.Find(&users)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get users: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User

	// select * from users where user_id = id
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(u *model.User) error {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&u).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	// delete from users where user_id = id
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		log.Printf("DeleteUser: Failed to delete user: %v\n", result.Error)
		return fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}

	return nil
}
