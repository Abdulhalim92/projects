package review

import "projects/internal/model"

type Service struct {
	ReviewRepository ReviewRepository
}

func NewService(r ReviewRepository) *Service {
	return &Service{r}
}

func (s *Service) AddReview(r *model.Reviews) (*model.Reviews, error) {
	return s.ReviewRepository.AddReview(r)
}

func (s *Service) DeleteReview(reviewID int) (int, error) {
	return s.ReviewRepository.DeleteReview(reviewID)
}

func (s *Service) GetReviewByID(reviewID int) (*model.Reviews, error) {
	return s.ReviewRepository.GetReviewByID(reviewID)
}

func (s *Service) GetReviewsByUser(userID int) ([]model.Reviews, error) {
	return s.ReviewRepository.GetReviewsByUser(userID)
}

func (s *Service) GetReviewsByBook(bookID int) ([]model.Reviews, error) {
	return s.ReviewRepository.GetReviewsByBook(bookID)
}

func (s *Service) UpdateReview(r *model.Reviews) (*model.Reviews, error) {
	return s.ReviewRepository.UpdateReview(r)
}

func (s *Service) GetReviews() ([]model.Reviews, error) {
	return s.ReviewRepository.GetReviews()
}
