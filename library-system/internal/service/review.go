package service

import (
	"errors"
	"fmt"
	"projects/internal/model"
)

func (s *Service) CreateReview(r *model.Reviews) (*model.Reviews, error) {
	// Проверяем существование пользователя по ID
	_, err := s.Repository.GetUserByID(r.UserID)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", r.UserID)
		}
		return nil, err
	}

	// Проверяем существование книги по ID
	_, err = s.Repository.GetBookByID(r.BookID)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book with id %d not found", r.BookID)
		}
		return nil, err
	}

	// Добавляем отзыв в репозиторий
	createdReview, err := s.Repository.AddReview(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
	}

	return createdReview, nil
}

func (s *Service) ListReviews() ([]model.Reviews, error) {
	// Получаем список отзывов из репозитория
	reviews, err := s.Repository.GetReviews()
	if err != nil {
		return nil, err
	}

	// Проверяем, что список отзывов не пустой
	if len(reviews) == 0 {
		return nil, fmt.Errorf("no reviews found")
	}

	return reviews, nil
}

func (s *Service) GetReviewByID(reviewID int) (*model.Reviews, error) {
	// Получаем отзыв из репозитория по указанному идентификатору
	reviewByID, err := s.Repository.GetReviewByID(reviewID)
	if err != nil {
		// Проверяем, является ли ошибка "record not found"
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("review with id %d not found", reviewID)
		}
		// Возвращаем любые другие ошибки, если они возникли
		return nil, err
	}

	// Возвращаем найденный отзыв
	return reviewByID, nil
}

func (s *Service) GetReviewsByUser(userID int) ([]model.Reviews, error) {
	// Проверяем существование пользователя по ID
	if _, err := s.Repository.GetUserByID(userID); err != nil {
		// Если ошибка — это "record not found", возвращаем сообщение о том, что пользователь не найден
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", userID)
		}
		// Возвращаем другие ошибки, если они возникли
		return nil, err
	}

	// Получаем список отзывов, оставленных пользователем с указанным идентификатором
	reviewsByUser, err := s.Repository.GetReviewsByUser(userID)
	if err != nil {
		return nil, err
	}

	// Если отзывов не найдено, возвращаем сообщение об этом
	if len(reviewsByUser) == 0 {
		return nil, fmt.Errorf("no reviews found for user with id %d", userID)
	}

	// Возвращаем список отзывов пользователя
	return reviewsByUser, nil
}

func (s *Service) GetReviewsByBook(bookID int) ([]model.Reviews, error) {
	// Проверяем существование книги по ID
	if _, err := s.Repository.GetBookByID(bookID); err != nil {
		// Если ошибка — это "record not found", возвращаем сообщение о том, что книга не найдена
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book with id %d not found", bookID)
		}
		// Возвращаем другие ошибки, если они возникли
		return nil, err
	}

	// Получаем список отзывов для книги с указанным идентификатором
	reviewsByBook, err := s.Repository.GetReviewsByBook(bookID)
	if err != nil {
		return nil, err
	}

	// Если отзывов не найдено, возвращаем сообщение об этом
	if len(reviewsByBook) == 0 {
		return nil, fmt.Errorf("no reviews found for book with id %d", bookID)
	}

	// Возвращаем список отзывов для книги
	return reviewsByBook, nil
}

func (s *Service) GetReviewsByFilter(filter model.ReviewFilter) ([]model.Reviews, error) {
	// Валидация фильтра
	if err := filter.ValidateReviewFilter(filter); err != nil {
		return nil, err
	}

	// Получаем список отзывов по фильтру из репозитория
	reviewsByFilter, err := s.Repository.GetReviewsByFilter(filter)
	if err != nil {
		return nil, err
	}

	// Если отзывов не найдено, возвращаем сообщение об этом
	if len(reviewsByFilter) == 0 {
		return nil, fmt.Errorf("no reviews found for the given filter")
	}

	// Возвращаем отфильтрованный список отзывов
	return reviewsByFilter, nil
}

func (s *Service) GetAverageRatingByFilter(filter model.ReviewFilter) (float64, error) {
	// Валидация фильтра
	if err := filter.ValidateReviewFilter(filter); err != nil {
		return 0, err
	}

	// Получаем средний рейтинг по фильтру из репозитория
	ratingByFilter, err := s.Repository.GetAverageRatingByFilter(filter)
	if err != nil {
		return 0, err
	}

	return ratingByFilter, nil
}

func (s *Service) EditReview(r *model.Reviews) (*model.Reviews, error) {
	// Проверяем, существует ли отзыв с указанным идентификатором
	_, err := s.Repository.GetReviewByID(r.ReviewID)
	if err != nil {
		// Если ошибка — это "record not found", возвращаем сообщение о том, что отзыв не найден
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("review with id %d not found", r.ReviewID)
		}
		// Возвращаем другие ошибки, если они возникли
		return nil, err
	}

	// Пытаемся обновить отзыв
	updatedReview, err := s.Repository.UpdateReview(r)
	if err != nil {
		return nil, fmt.Errorf("failed to update review with id %d: %w", r.ReviewID, err)
	}

	return updatedReview, nil
}

func (s *Service) DeleteReview(reviewID int) (int, error) {
	// Проверяем, существует ли отзыв с указанным идентификатором
	_, err := s.Repository.GetReviewByID(reviewID)
	if err != nil {
		// Если ошибка — это "record not found", возвращаем сообщение о том, что отзыв не найден
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("review with id %d not found", reviewID)
		}
		// Возвращаем другие ошибки, если они возникли
		return 0, err
	}

	// Пытаемся удалить отзыв
	deletedRows, err := s.Repository.DeleteReview(reviewID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete review with id %d: %w", reviewID, err)
	}

	return deletedRows, nil
}
