package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) AddReview(review *model.Reviews) (*model.Reviews, error) {
	// insert into reviews (book_id, user_id, rating, comment) values (1, 1, 5, 'Good')
	result := r.db.Create(&review)
	if result.Error != nil {
		log.Printf("CreateReview: Failed to add review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add review: %v\n", result.Error)
	}

	return review, nil

}

func (r *Repository) GetReviews() ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews
	err := r.db.Find(&reviews).Error
	if err != nil {
		log.Printf("ListReviews: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("Cannot find reviews with error: %v", err)
	}

	return reviews, nil
}

func (r *Repository) GetReviewsByFilter(filter model.ReviewFilter) ([]model.Reviews, error) {
	var reviews []model.Reviews

	query := r.db

	if filter.ReviewID > 0 {
		query = query.Where("review_id = ?", filter.ReviewID)
	}
	if filter.BookID > 0 {
		query = query.Where("book_id = ?", filter.BookID)
	}

	// select * from reviews where book_id = bookID
	err := r.db.Where(filter).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByFilter: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("cannot find reviews with error: %v", err)
	}

	return reviews, nil
}

func (r *Repository) GetAverageRatingByFilter(filter model.ReviewFilter) (float64, error) {
	var average float64

	// select avg(rating) from reviews
	err := r.db.Table("reviews").Select("avg(rating)").Where(filter).Find(&average).Error
	if err != nil {
		log.Printf("GetAverageRatingByFilter: Failed to get average rating: %v\n", err)
		return 0, fmt.Errorf("cannot find average rating with error: %v", err)
	}

	return average, nil
}

func (r *Repository) DeleteReview(reviewID int) (int, error) {
	// delete from reviews where review_id = reviewID
	result := r.db.Delete(&model.Reviews{}, reviewID)
	if result.Error != nil {
		log.Printf("DeleteReview: Failed to delete review: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete review: %v\n", result.Error)
	}

	return reviewID, nil

}

func (r *Repository) GetReviewByID(reviewID int) (*model.Reviews, error) {
	var review model.Reviews

	// select * from reviews where review_id = reviewID
	err := r.db.Where("review_id = ?", reviewID).Find(&review).Error
	if err != nil {
		log.Printf("GetReviewByID: Failed to get review: %v\n", err)
		return nil, fmt.Errorf("Cannot find review with error: %v\n", err)
	}

	return &review, nil
}

func (r *Repository) GetReviewsByUser(userID int) ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews where user_id = userID
	err := r.db.Where("user_id = ?", userID).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByUser: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("cannot find reviews with error: %v", err)
	}

	return reviews, nil

}

func (r *Repository) GetReviewsByBook(bookID int) ([]model.Reviews, error) {
	var reviews []model.Reviews

	// select * from reviews where book_id = bookID
	err := r.db.Where("book_id = ?", bookID).Find(&reviews).Error
	if err != nil {
		log.Printf("GetReviewsByBook: Failed to get reviews: %v\n", err)
		return nil, fmt.Errorf("cannot find reviews with error: %v", err)
	}

	return reviews, nil

}

func (r *Repository) UpdateReview(review *model.Reviews) (*model.Reviews, error) {
	// update reviews set rating = 5, comment = 'Good' where review_id = 1
	result := r.db.Save(&review)
	if result.Error != nil {
		log.Printf("EditReview: Failed to update review: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update review: %v\n", result.Error)
	}

	return review, nil
}
