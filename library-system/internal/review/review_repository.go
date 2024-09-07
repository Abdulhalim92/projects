package review

import (
	"fmt"
	"log"
	"projects/internal/model"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) AddReview(review *model.Reviews) (*model.Reviews, error) {
	// insert into reviews (book_id, user_id, rating, comment) values (1, 1, 5, 'Good')
	result := r.db.Create(&review)
	if result.Error != nil {
		log.Printf("AddReview: Failed to add review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add review: %v\n", result.Error)
	}

	return review, nil

}

func (r *ReviewRepository) GetReviews() ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews
	err := r.db.Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviews: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("Cannot find reviews with error: %v", err)
	}

	return reviews, nil
}

func (r *ReviewRepository) DeleteReview(reviewID int) (int, error) {
	// delete from reviews where review_id = reviewID
	result := r.db.Delete(&model.Reviews{}, reviewID)
	if result.Error != nil {
		log.Printf("DeleteReview: Failed to delete review: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete review: %v\n", result.Error)
	}

	return reviewID, nil

}

func (r *ReviewRepository) GetReviewByID(reviewID int) (*model.Reviews, error) {
	var review model.Reviews

	// select * from reviews where review_id = reviewID
	err := r.db.Where("review_id = ?", reviewID).Find(&review).Error
	if err != nil {
		log.Printf("GetReviewByID: Failed to get review: %v\n", err)
		return nil, fmt.Errorf("Cannot find review with error: %v", err)
	}

	return &review, nil
}

func (r *ReviewRepository) GetReviewsByUser(userID int) ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews where user_id = userID
	err := r.db.Where("user_id = ?", userID).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByUser: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("Cannot find reviews with error: %v", err)
	}

	return reviews, nil

}

func (r *ReviewRepository) GetReviewsByBook(bookID int) ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews where book_id = bookID
	err := r.db.Where("book_id = ?", bookID).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByBook: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("Cannot find reviews with error: %v", err)
	}

	return reviews, nil

}

func (r *ReviewRepository) UpdateReview(review *model.Reviews) (*model.Reviews, error) {
	// update reviews set rating = 5, comment = 'Good' where review_id = 1
	result := r.db.Save(&review)
	if result.Error != nil {
		log.Printf("UpdateReview: Failed to update review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update review: %v\n", result.Error)
	}

	return review, nil
}
