package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"projects/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRoles() ([]model.Role, error) {
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

func (r *Repository) GetRoleByID(roleID int) (*model.Role, error) {
	var role *model.Role
	// select * from roles where role_id = ?
	err := r.db.First(&role, roleID).Error
	if err != nil {
		log.Printf("GetRoleByID: Error getting role by ID: %v", err)
		return nil, err
	}

	return role, nil
}

func (r *Repository) CreateRole(role *model.Role) (int, error) {
	// insert into roles (name) values ('admin') returning role_id
	err := r.db.Create(role).Error
	if err != nil {
		log.Printf("CreateROle: Error creating role: %v", err)
		return 0, err
	}

	return role.RoleID, nil
}

func (r *Repository) UpdateRole(role *model.Role) (*model.Role, error) {
	// update roles set name = 'admin' where role_id = 1
	err := r.db.Clauses(clause.Returning{}).Updates(role).Error
	if err != nil {
		log.Printf("UpdateRole: Error updating role: %v", err)
		return nil, err
	}

	return role, nil
}

func (r *Repository) DeleteRole(roleID int) error {
	// delete from roles where role_id = 1
	err := r.db.Delete(&model.Role{}, roleID).Error
	if err != nil {
		log.Printf("DeleteRole: Error deleting role: %v", err)
		return err
	}

	return nil
}
