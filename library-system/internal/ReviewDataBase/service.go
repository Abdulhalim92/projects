package ReviewDataBase

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	r *ReviewRep
}

func NewService(r *ReviewRep) *Service {
	return &Service{r: r}
}
func (s *Service) CreateReview(r *model.Review) (*model.Review, error) {
	return s.r.AddReview(r)
}
func (s *Service) ListReviews() ([]model.Review, error) {
	reviews, err := s.r.GetReviews()
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews exists")
	}
	return reviews, nil
}
func (s *Service) ListReviewById(id int) (*model.Review, error) {
	review, err := s.r.GetReviewById(id)
	if err != nil {
		return nil, err
	} else if review.ReviewID == 0 {
		return nil, fmt.Errorf("no review with such id")
	}
	return review, nil
}
func (s *Service) ListReviewsByUserId(UserID int) ([]model.Review, error) {
	reviews, err := s.r.GetReviewsByUserID(UserID)
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews with such userID")
	}
	return reviews, nil
}
func (s *Service) ListReviewsByBookId(BookID int) ([]model.Review, error) {
	reviews, err := s.r.GetReviewsByBookID(BookID)
	if err != nil {
		return nil, err
	} else if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews with such bookID")
	}
	return reviews, nil
}
func (s *Service) EditReview(rev *model.Review) (*model.Review, error) {
	review, err := s.r.GetReviewById(rev.ReviewID)
	if err != nil {
		return nil, err
	} else if review.ReviewID == 0 {
		return nil, fmt.Errorf("no review with such id")
	}
	return s.r.UpdateReview(rev)
}
func (s *Service) RemoveReviewById(id int) (int, error) {
	review, err := s.r.GetReviewById(id)
	if err != nil {
		return 0, err
	} else if review.ReviewID == 0 {
		return 0, fmt.Errorf("such review doesn't exist")
	}
	return s.r.DeleteReviewById(id)
}
