package repository

import (
	"fmt"
	"projects/internal/model"
)

func (r *Repository) GetReviews() ([]*model.Review, error) {
	var reviews []*model.Review
	result := r.db.Find(&reviews)
	if result.Error != nil {
		fmt.Errorf("Failed to get reviews %v\n", result.Error)
		return nil, result.Error
	}
	return reviews, nil
}

func (r *Repository) GetReviewsByFilter(filter *model.ReviewFilter) ([]*model.Review, error) {
	var reviews []*model.Review
	query := r.db
	if filter.ReviewID > 0 {
		query = query.Where("review_id = ?", filter.ReviewID).Find(&reviews)
	}
	if filter.BookID > 0 {
		query = query.Where("book_id = ?", filter.BookID).Find(&reviews)
	}
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID).Find(&reviews)
	}
	if !filter.DateFrom.IsZero() {
		query = query.Where("created_at >= ?", filter.DateFrom).Find(&reviews)
	}
	if !filter.DateTo.IsZero() {
		query = query.Where("created_at <= ?", filter.DateTo).Find(&reviews)
	}
	query.Offset((filter.Page - 1) * filter.CountOnPage).Limit(filter.CountOnPage).Find(&reviews)

	if query.Error != nil {
		fmt.Errorf("Failed to get filtered borrows")
		return nil, query.Error
	}
	return reviews, nil
}

func (r *Repository) AddReview(review *model.Review) (*model.Review, error) {
	result := r.db.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}
