package service

import (
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateReview(r *model.Reviews) (*model.Reviews, error) {
	userByID, err := s.Repository.GetUserByID(r.UserID)
	if err != nil {
		return nil, err
	}
	if userByID == nil {
		return nil, fmt.Errorf("user with id %d not found", r.UserID)
	}

	bookByID, err := s.Repository.GetBookByID(r.BookID)
	if err != nil {
		return nil, err
	}
	if bookByID == nil {
		return nil, fmt.Errorf("book with id %d not found", r.BookID)
	}

	return s.Repository.AddReview(r)
}

func (s *Service) ListReviews() ([]model.Reviews, error) {
	reviews, err := s.Repository.GetReviews()
	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews found")
	}

	return reviews, nil
}

func (s *Service) GetReviewByID(reviewID int) (*model.Reviews, error) {
	reviewByID, err := s.Repository.GetReviewByID(reviewID)
	if err != nil {
		return nil, err
	}

	if reviewByID == nil {
		return nil, fmt.Errorf("review with id %d not found", reviewID)
	}

	return reviewByID, nil
}

func (s *Service) GetReviewsByUser(userID int) ([]model.Reviews, error) {
	reviewsByUser, err := s.Repository.GetReviewsByUser(userID)
	if err != nil {
		return nil, err
	}

	if len(reviewsByUser) == 0 {
		return nil, fmt.Errorf("no reviews found for user with id %d", userID)
	}

	return reviewsByUser, nil
}

func (s *Service) GetReviewsByBook(bookID int) ([]model.Reviews, error) {
	reviewsByBook, err := s.Repository.GetReviewsByBook(bookID)
	if err != nil {
		return nil, err
	}

	if len(reviewsByBook) == 0 {
		return nil, fmt.Errorf("no reviews found for book with id %d", bookID)
	}

	return reviewsByBook, nil
}

func (s *Service) GetReviewsByFilter(filter model.ReviewFilter) ([]model.Reviews, error) {
	reviewsByFilter, err := s.Repository.GetReviewsByFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(reviewsByFilter) == 0 {
		return nil, fmt.Errorf("no reviews found")
	}

	return reviewsByFilter, nil
}

func (s *Service) GetAverageRatingByFilter(filter model.ReviewFilter) (float64, error) {
	ratingByFilter, err := s.Repository.GetAverageRatingByFilter(filter)
	if err != nil {
		return 0, err
	}

	return ratingByFilter, nil
}

func (s *Service) EditReview(r *model.Reviews) (*model.Reviews, error) {
	reviewByID, err := s.Repository.GetReviewByID(r.ReviewID)
	if err != nil {
		return nil, err
	}

	if reviewByID == nil {
		return nil, fmt.Errorf("review with id %d not found", r.ReviewID)
	}

	return s.Repository.UpdateReview(r)
}

func (s *Service) DeleteReview(reviewID int) (int, error) {
	reviewByID, err := s.Repository.GetReviewByID(reviewID)
	if err != nil {
		return 0, err
	}

	if reviewByID == nil {
		return 0, fmt.Errorf("review with id %d not found", reviewID)
	}

	return s.Repository.DeleteReview(reviewID)
}
