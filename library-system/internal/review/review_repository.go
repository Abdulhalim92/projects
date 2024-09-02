package review

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) AddReview(rv *model.Review) (*model.Review, error) {
	// insert into reviews (reviewname, password) values ('admin', 'admin')
	result := r.db.Create(&rv)
	if result.Error != nil {
		log.Printf("AddReview: Failed to add review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add review: %v\n", result.Error)
	}

	return rv, nil
}

func (r *ReviewRepo) GetReviews() ([]model.Review, error) {
	var reviews []model.Review

	// select * from reviews
	result := r.db.Find(&reviews)
	if result.Error != nil {
		log.Printf("GetReviews: Failed to get reviews: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get reviews: %v\n", result.Error)
	}
	return reviews, nil
}

func (r *ReviewRepo) GetReviewByID(id int) (*model.Review, error) {
	var review model.Review

	// select * from reviews where review_id = id
	result := r.db.First(&review, id)
	if result.Error != nil {
		log.Printf("GetReviewByID: Failed to get review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get review: %v\n", result.Error)
	}
	return &review, nil
}

func (r *ReviewRepo) GetReviewsByUserID(userID int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := r.db.Find(&borrows, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return borrows, nil
}

func (r *ReviewRepo) GetReviewsByBookID(bookID int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := r.db.Find(&borrows, "book_id = ?", bookID).Error
	if err != nil {
		return nil, err
	}

	return borrows, nil
}

func (r *ReviewRepo) UpdateReview(u *model.Review) error {
	// update reviews set reviewname = 'admin', password = 'admin' where review_id = 1
	result := r.db.Model(&u).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateReview: Failed to update review: %v\n", result.Error)
		return fmt.Errorf("Failed to update review: %v\n", result.Error)
	}

	return nil
}

func (r *ReviewRepo) DeleteReview(id int) error {
	// delete from reviews where review_id = id
	result := r.db.Delete(&model.Review{}, id)
	if result.Error != nil {
		log.Printf("DeleteReview: Failed to delete review: %v\n", result.Error)
		return fmt.Errorf("Failed to delete review: %v\n", result.Error)
	}

	return nil
}
