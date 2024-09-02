package ReviewDataBase

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type ReviewRep struct {
	db *gorm.DB
}

func NewReviewRep(db *gorm.DB) *ReviewRep {
	return &ReviewRep{db: db}
}
func (r *ReviewRep) AddReview(review *model.Review) (*model.Review, error) {
	err := r.db.Table("reviews").Create(review).Error
	if err != nil {
		return nil, fmt.Errorf("error adding a review: %v", err)
	}
	return review, nil
}
func (r *ReviewRep) GetReviews() ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Find(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error getting reviews: %v", err)
	}
	return reviews, nil
}
func (r *ReviewRep) GetReviewById(id int) (*model.Review, error) {
	var review model.Review
	err := r.db.Table("reviews").Where("review_id = ?", id).Scan(&review).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a review by id: %v", err)
	}
	return &review, err
}
func (r *ReviewRep) UpdateReview(review *model.Review) (*model.Review, error) {
	err := r.db.Updates(review).Error
	if err != nil {
		return nil, fmt.Errorf("error updating a review: %v", err)
	}
	return review, nil
}
func (r *ReviewRep) DeleteReviewById(id int) (int, error) {
	err := r.db.Delete(&model.Review{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting a review: %v", err)
	}
	return id, nil
}
func (r *ReviewRep) GetReviewsByBookID(BookID int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Where("book_id = ?", BookID).Scan(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error getting reviews by bookID: %v", err)
	}
	return reviews, nil
}
func (r *ReviewRep) GetReviewsByUserID(UserID int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Where("user_id = ?", UserID).Scan(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting reviews by userID: %v", err)
	}
	return reviews, nil
}
