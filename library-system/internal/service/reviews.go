package service

import "projects/internal/model"

func (s *Service) ListReviews() ([]*model.Review, error) {
	return s.Repository.GetReviews()
}

func (s *Service) CreateReview(review *model.Review) (*model.Review, error) {
	return s.Repository.AddReview(review)
}

func (s *Service) ListFilteredReviews(filter *model.ReviewFilter) ([]*model.Review, error) {
	return s.Repository.GetReviewsByFilter(filter)
}
