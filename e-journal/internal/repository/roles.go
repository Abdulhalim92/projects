package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"projects/internal/model"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetRoles() ([]model.Role, error) {
	var roles []model.Role
	// select * from roles
	err := r.db.Find(&roles).Error
	if err != nil {
		log.Printf("GetRoles: Error getting roles: %v", err)
		return nil, err
	}
	if len(roles) == 0 {
		log.Printf("GetRoles: No roles found")
		return nil, fmt.Errorf("GetRoles: No roles found")
	}

	return roles, nil
}

func (r *RoleRepository) GetRoleByID(roleID int) (*model.Role, error) {
	var role *model.Role
	// select * from roles where role_id = ?
	err := r.db.First(&role, roleID).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("record not found")
		}
		log.Printf("GetRoleByID: Error getting role by ID: %v", err)
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) CreateRole(role *model.Role) (int, error) {
	// insert into roles (name) values ('admin') returning role_id
	err := r.db.Create(role).Error
	if err != nil {
		log.Printf("CreateROle: Error creating role: %v", err)
		return 0, err
	}

	return role.RoleID, nil
}
