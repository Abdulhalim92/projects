package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) GetParents() ([]model.Parent, error) {
	var parents []model.Parent
	// select * from parents
	err := r.db.Find(&parents).Error
	if err != nil {
		log.Printf("GetParents: Error getting parents: %v", err)
		return nil, err
	}
	if len(parents) == 0 {
		log.Printf("GetParents: No parents found")
		return nil, fmt.Errorf("GetParents: No parents found")
	}

	return parents, nil
}

func (r *Repository) CreateParent(parent *model.Parent) (int, error) {
	// insert into parents (name, user_id, student_id) values ('admin', 1, 1) returning parent_id
	err := r.db.Create(parent).Error
	if err != nil {
		log.Printf("CreateParent: Error creating parent: %v", err)
	}
	return parent.ParentID, nil
}

func (r *Repository) GetParentByID(parentID int) (*model.Parent, error) {
	var parent *model.Parent
	// select * from parents where parent_id = ?
	err := r.db.First(&parent, parentID).Error
	if err != nil {
		log.Printf("GetParentByID: Error getting parent by ID: %v", err)
		return nil, err
	}
	return parent, nil
}
func (r *Repository) UpdateParent(parent *model.Parent) (*model.Parent, error) {
	//update parents sen name ='admin ' where parent_id = 1
	err := r.db.Updates(parent).Error
	if err != nil {
		log.Printf("UpdateParent: Error updating parent: %v", err)
		return nil, err
	}
	return parent, nil
}
func (r *Repository) DeleteParent(parentID int) error {
	// delete from parents where parent_id = 1
	err := r.db.Delete(&model.Parent{}, parentID).Error
	if err != nil {
		log.Printf("DeleteParent: Error deleting parent: %v", err)
		return err
	}
	return nil

}
