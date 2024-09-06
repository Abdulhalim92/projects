package profile

import (
	"fmt"
	"library-system/internal/model"
)

type Service struct {
	pr ProfileRepo
}

func NewService(pr ProfileRepo) *Service {
	return &Service{pr}
}

func (s *Service) CreateProfile(userID int, email, address string) (*model.Profile, error) {
	profile := model.Profile{UserId: userID, Email: email, Address: address}
	return s.pr.AddProfile(&profile)
}

func (s *Service) ListProfiles() ([]model.Profile, error) {
	profiles, err := s.pr.GetProfiles()
	if err != nil {

		return nil, fmt.Errorf("Error when listing the profiles: %e", err)
	}

	return profiles, nil
}

func (s *Service) FindProfile(userID int) (*model.Profile, error) {
	profile, err := s.pr.GetProfileByUserID(userID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error occured when retrieiving profile with user_id:%d\n%e", userID, err)
	}

	return profile, nil
}

func (s *Service) EditProfile(userID int, email, address string) error {
	profile, err := s.FindProfileByEmail(email)
	if err != nil {
		return err
	}

	err = s.pr.UpdateProfile(profile)
	if err != nil {
		return fmt.Errorf("Error occured when editing profile with user_id:%d\n%e", userID, err)
	}
	return nil
}

func (s *Service) RemoveProfile(id int) bool {
	err := s.pr.DeleteProfile(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) FindProfileByEmail(email string) (*model.Profile, error) {
	var profile model.Profile
	err := s.pr.db.First(&profile, "email = ?", email).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't find profile with profilename:%s\n%e", email, err)
	}

	return &profile, nil
}
