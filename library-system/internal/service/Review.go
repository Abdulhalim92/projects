package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateReview(r *model.Review) (*model.Review, error) {
	return s.Repository.AddReview(r)
}
func (s *Service) ListReviews() ([]model.Review, error) {
	reviews, err := s.Repository.GetReviews()
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews exists")
	}
	return reviews, nil
}
func (s *Service) ListReviewById(id int) (*model.Review, error) {
	review, err := s.Repository.GetReviewById(id)
	if err != nil {
		return nil, err
	} else if review.ReviewID == 0 {
		return nil, fmt.Errorf("no review with such id")
	}
	return review, nil
}
func (s *Service) ListReviewsByUserId(UserID int) ([]model.Review, error) {
	reviews, err := s.Repository.GetReviewsByUserID(UserID)
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews with such userID")
	}
	return reviews, nil
}
func (s *Service) ListReviewsByBookId(BookID int) ([]model.Review, error) {
	reviews, err := s.Repository.GetReviewsByBookID(BookID)
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews with such bookID")
	}
	return reviews, nil
}
func (s *Service) EditReview(rev *model.Review) (*model.Review, error) {
	review, err := s.Repository.GetReviewById(rev.ReviewID)
	if err != nil {
		return nil, err
	} else if review.ReviewID == 0 {
		return nil, fmt.Errorf("no review with such id")
	}
	return s.Repository.UpdateReview(rev)
}
func (s *Service) RemoveReviewById(id int) (int, error) {
	review, err := s.Repository.GetReviewById(id)
	if err != nil {
		return 0, err
	} else if review.ReviewID == 0 {
		return 0, fmt.Errorf("such review doesn't exist")
	}
	return s.Repository.DeleteReviewById(id)
}
func (s *Service) GetReviewsByFilter(filter model.ReviewFilter) ([]model.Review, error) {
	reviewsByFilter, err := s.Repository.GetReviewsByFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(reviewsByFilter) == 0 {
		return nil, fmt.Errorf("no reviews found")
	}

	return reviewsByFilter, nil
}
