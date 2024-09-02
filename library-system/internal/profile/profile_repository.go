package profile

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfileRepo(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (r *ProfileRepo) AddProfile(p *model.Profile) (*model.Profile, error) {
	// insert into profiles (profilename, password) values ('admin', 'admin')
	result := r.db.Create(&p)
	if result.Error != nil {
		log.Printf("AddProfile: Failed to add profile: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add profile: %v\n", result.Error)
	}

	return p, nil
}

func (r *ProfileRepo) GetProfiles() ([]model.Profile, error) {
	var profiles []model.Profile

	// select * from profiles
	result := r.db.Find(&profiles)
	if result.Error != nil {
		log.Printf("GetProfiles: Failed to get profiles: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get profiles: %v\n", result.Error)
	}
	return profiles, nil
}

func (r *ProfileRepo) GetProfileByUserID(userID int) (*model.Profile, error) {
	var profile model.Profile

	result := r.db.First(&profile, userID)
	if result.Error != nil {
		log.Printf("GetProfileByID: Failed to get profile: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get profile: %v\n", result.Error)
	}
	return &profile, nil
}

func (r *ProfileRepo) UpdateProfile(u *model.Profile) error {
	result := r.db.Model(&u).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateProfile: Failed to update profile: %v\n", result.Error)
		return fmt.Errorf("Failed to update profile: %v\n", result.Error)
	}

	return nil
}

func (r *ProfileRepo) DeleteProfile(id int) error {
	result := r.db.Delete(&model.Profile{}, id)
	if result.Error != nil {
		log.Printf("DeleteProfile: Failed to delete profile: %v\n", result.Error)
		return fmt.Errorf("Failed to delete profile: %v\n", result.Error)
	}

	return nil
}
