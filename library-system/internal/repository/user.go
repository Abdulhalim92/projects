package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"projects/internal/model"
)

func (r *Repository) AddUser(u *model.User) (*model.User, error) {
	// insert into users (username, password) values ('admin', 'admin')
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Printf("SignUp: Failed to add user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}

	return u, nil
}

func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User

	// select * from users
	result := r.db.Omit("password").Find(&users)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get users: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*model.User, error) {
	var user model.User

	// select * from users where user_id = id
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	// select * from users where username = username
	result := r.db.First(&user, "username = ?", username)
	if result.Error != nil {
		log.Printf("GetUserByUsername: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *Repository) UpdateUser(u *model.User) (*model.User, error) {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&u).Clauses(clause.Returning{}).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return u, nil
}

func (r *Repository) DeleteUser(id int) (int, error) {
	// delete from users where user_id = id
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		log.Printf("DeleteUser: Failed to delete user: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}

	return id, nil
}
