package model

import (
	"fmt"
	"time"
)

type Book struct {
	BookID    int        `json:"book_id,omitempty" gorm:"primaryKey"`
	Title     string     `json:"title,omitempty"`
	AuthorID  int        `json:"author_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Reviews struct {
	ReviewID   int        `json:"review_id" gorm:"primaryKey"`
	UserID     int        `json:"user_id"`
	BookID     int        `json:"book_id"`
	ReviewText string     `json:"review_text"`
	Rating     float64    `json:"rating"`
	ReviewDate *time.Time `json:"review_date,omitempty"`
}

type Borrow struct {
	BorrowID   int        `json:"borrow_id" gorm:"primaryKey"`
	UserID     int        `json:"user_id"`
	BookID     int        `json:"book_id"`
	BorrowDate *time.Time `json:"borrow_date,omitempty"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
}

type BorrowHistory struct {
	HistoryID  int        `json:"history_id" gorm:"primaryKey"`
	BorrowID   int        `json:"borrow_id"`
	ActionType string     `json:"action_type"`
	ActionDate *time.Time `json:"action_date,omitempty"`
}

type ReviewFilter struct {
	ReviewID    int        `json:"review_id,omitempty"`
	BookID      int        `json:"book_id,omitempty"`
	UserID      int        `json:"user_id,omitempty"`
	CountOnPage int        `json:"count_on_page,omitempty"`
	Page        int        `json:"page,omitempty"`
	DateFrom    *time.Time `json:"date_from,omitempty"`
	DateTo      *time.Time `json:"date_to,omitempty"`
}

// ValidateReviewFilter проверяет корректность значений фильтра
func (r *ReviewFilter) ValidateReviewFilter(filter ReviewFilter) error {
	if filter.CountOnPage < 0 {
		return fmt.Errorf("CountOnPage must be non-negative")
	}
	if filter.Page < 0 {
		return fmt.Errorf("page must be non-negative")
	}
	if filter.DateFrom != nil && filter.DateTo != nil && filter.DateFrom.After(*filter.DateTo) {
		return fmt.Errorf("DateFrom cannot be after DateTo")
	}
	return nil
}
