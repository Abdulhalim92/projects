package review

import (
	"fmt"
	"library-system/internal/model"
	"time"
)

type Service struct {
	rr ReviewRepo
}

func NewService(rr ReviewRepo) *Service {
	return &Service{rr}
}

func (s *Service) CreateReview(userID, bookID int, reviewText string, reviewDate time.Time) (*model.Review, error) {
	review := model.Review{UserId: userID, BookId: bookID, ReviewText: reviewText, ReviewDate: reviewDate}
	return s.rr.AddReview(&review)
}

func (s *Service) ListReviews() ([]model.Review, error) {
	reviews, err := s.rr.GetReviews()
	if err != nil {

		return nil, fmt.Errorf("Error when listing the reviews: %e", err)
	}

	return reviews, nil
}

func (s *Service) FindReview(id int) (*model.Review, error) {
	review, err := s.rr.GetReviewByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error occured when retrieiving review with id:%d\n%e", id, err)
	}

	return review, nil
}

func (s *Service) FindReviewsByBookTitle(title string) ([]model.Review, error) {
	var reviews []model.Review
	err := s.rr.db.Table("reviews").Select("reviews").Joins("JOIN books ON reviews.book_id = books.id").Where("book.title = ?", title).Scan(&reviews)
	if len(reviews) == 0 {
		return nil, fmt.Errorf("Couldn't get reviews of book with title:%d\n%e", title, err)
	}

	return reviews, nil
}

func (s *Service) EditReview(id, user_id, book_id int, reviewText string, rating float32, reviewDate time.Time) error {
	review, err := s.FindReview(id)
	if err != nil {
		return err
	}

	err = s.rr.UpdateReview(review)
	if err != nil {
		return fmt.Errorf("Error occured when editing review with id:%d\n%e", id, err)
	}
	return nil
}

func (s *Service) RemoveReview(id int) bool {
	err := s.rr.DeleteReview(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
