package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"projects/internal/model"
)

func (r *Repository) GetAdmins() ([]model.Admin, error) {
	var admins []model.Admin
	// select * from admins
	err := r.db.Find(&admins).Error
	if err != nil {
		log.Printf("GetAdmins: Error getting admins: %v", err)
		return nil, err
	}
	if len(admins) == 0 {
		log.Printf("GetAdmins: No admins found")
		return nil, fmt.Errorf("GetAdmins: No admins found")
	}

	return admins, nil

}
func (r *Repository) CreateAdmin(admin *model.Admin) (int, error) {
	// insert into admins (name, user_id, student_id) values ('admin', 1, 1) returning admin_id
	err := r.db.Create(admin).Error
	if err != nil {
		log.Printf("CreateAdmin: Error creating admin: %v", err)
	}
	return admin.AdminID, nil

}

func (r *Repository) UpdateAdmin(admin *model.Admin) (*model.Admin, error) {
	// update admins set name = 'admin' where admin_id = 1
	err := r.db.Clauses(clause.Returning{}).Updates(admin).Error
	if err != nil {
		log.Printf("UpdateAdmin: Error updating admin: %v", err)
		return nil, err
	}

	return admin, nil
}

func (r *Repository) DeleteAdmin(adminID int) error {
	// delete from admins where admin_id = 1
	err := r.db.Delete(&model.Admin{}, adminID).Error
	if err != nil {
		log.Printf("DeleteAdmin: Error deleting admin: %v", err)
		return err

	}
	return nil
}

func (r *Repository) GetAdminByID(adminID int) (*model.Admin, error) {
	var admin *model.Admin
	// select * from admins where admin_id = ?
	err := r.db.First(&admin, adminID).Error
	if err != nil {
		log.Printf("GetAdminByID: Error getting admin by ID: %v", err)
		return nil, err
	}
	return admin, nil

}
