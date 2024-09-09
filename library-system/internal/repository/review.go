package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) AddReview(review *model.Review) (*model.Review, error) {
	err := r.db.Table("reviews").Create(review).Error
	if err != nil {
		return nil, fmt.Errorf("error adding a review: %v", err)
	}
	return review, nil
}
func (r *Repository) GetReviews() ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Find(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error getting reviews: %v", err)
	}
	return reviews, nil
}
func (r *Repository) GetReviewById(id int) (*model.Review, error) {
	var review model.Review
	err := r.db.Table("reviews").Where("review_id = ?", id).Scan(&review).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a review by id: %v", err)
	}
	return &review, err
}
func (r *Repository) UpdateReview(review *model.Review) (*model.Review, error) {
	err := r.db.Updates(review).Error
	if err != nil {
		return nil, fmt.Errorf("error updating a review: %v", err)
	}
	return review, nil
}
func (r *Repository) DeleteReviewById(id int) (int, error) {
	err := r.db.Delete(&model.Review{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting a review: %v", err)
	}
	return id, nil
}
func (r *Repository) GetReviewsByBookID(BookID int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Where("book_id = ?", BookID).Scan(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error getting reviews by bookID: %v", err)
	}
	return reviews, nil
}
func (r *Repository) GetReviewsByUserID(UserID int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Table("reviews").Where("user_id = ?", UserID).Scan(&reviews).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting reviews by userID: %v", err)
	}
	return reviews, nil
}
func (r *Repository) GetReviewsByFilter(filter model.ReviewFilter) ([]model.Review, error) {
	var reviews []model.Review

	// select * from reviews where book_id = bookID
	err := r.db.Where(filter).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByFilter: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("cannot find reviews with error: %v", err)
	}

	return reviews, nil
}
