package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) ListProfiles() ([]model.Profile, error) {
	profiles, err := s.Repository.GetProfiles()
	if err != nil {
		return nil, err
	}

	if len(profiles) == 0 {
		return nil, fmt.Errorf("no profiles found")
	}

	return profiles, nil
}

func (s *Service) CreateProfile(p *model.Profile) (*model.Profile, error) {
	userByID, err := s.Repository.GetUserByID(p.UserID)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", p.UserID)
	}

	profileByID, err := s.Repository.GetProfileByID(p.UserID)
	if err != nil {
		return nil, err
	}
	if profileByID != nil {
		return nil, fmt.Errorf("profile with user id %d already exists", p.UserID)
	}

	return s.Repository.AddProfile(p)
}

func (s *Service) EditProfile(p *model.Profile) (*model.Profile, error) {
	userByID, err := s.Repository.GetUserByID(p.UserID)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", p.UserID)
	}

	profileByID, err := s.Repository.GetProfileByID(p.UserID)
	if err != nil {
		return nil, err
	}
	if profileByID == nil {
		return nil, fmt.Errorf("profile with user id %d not found", p.UserID)
	}

	return s.Repository.UpdateProfile(p)
}

func (s *Service) GetProfileByID(id int) (*model.Profile, error) {
	profileByID, err := s.Repository.GetProfileByID(id)
	if err != nil {
		return nil, err
	}

	if profileByID == nil {
		return nil, fmt.Errorf("profile with id %d not found", id)
	}

	return profileByID, nil
}

func (s *Service) DeleteProfile(id int) (int, error) {
	profileByID, err := s.Repository.GetProfileByID(id)
	if err != nil {
		return 0, err
	}

	if profileByID == nil {
		return 0, fmt.Errorf("profile with id %d not found", id)
	}

	return s.Repository.DeleteProfile(id)
}
