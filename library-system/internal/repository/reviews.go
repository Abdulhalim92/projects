package repository

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type ReviewsRepo struct {
	db *gorm.DB
}

func NewReviewsRepo(db *gorm.DB) *ReviewsRepo {
	return &ReviewsRepo{db: db}
}

func (r *ReviewsRepo) GetReviews() ([]model.Review, error) {
	var reviews []model.Review
	result := r.db.Find(&reviews)
	if result.Error != nil {
		fmt.Errorf("Failed to get reviews %v\n", result.Error)
		return nil, result.Error
	}
	return reviews, nil
}

func (r *ReviewsRepo) GetReviewsByFilter(filter *model.ReviewFilter) ([]model.Review, error) {

	var reviews []model.Review
	query := r.db

	if filter.ReviewID > 0 {
		query = query.Where("review_id = ?", filter.ReviewID).Find(&reviews)
	}

	if filter.BookID > 0 {
		query = query.Where("book_id = ?", filter.BookID).Find(&reviews)
	}

	query.Offset((filter.Page - 1) * filter.CountOnPage).Limit(filter.CountOnPage)

	return reviews, nil

}

func (r *ReviewsRepo) AddReview(review *model.Review) (*model.Review, error) {
	result := r.db.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}
