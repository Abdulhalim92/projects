package model

import "time"

type Book struct {
	BookID    int       `json:"book_id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reviews struct {
	ReviewID   int     `json:"review_id" gorm:"primaryKey"`
	UserID     int     `json:"user_id"`
	BookID     int     `json:"book_id"`
	ReviewText string  `json:"review_text"`
	Rating     float64 `json:"rating"`
	ReviewDate *time.Time
}

type Borrow struct {
	BorrowID   int `json:"borrow_id" gorm:"primaryKey"`
	UserID     int `json:"user_id"`
	BookID     int `json:"book_id"`
	BorrowDate *time.Time
	ReturnDate *time.Time
}

type BorrowHistory struct {
	HistoryID  int    `json:"history_id" gorm:"primaryKey"`
	BorrowID   int    `json:"borrow_id"`
	ActionType string `json:"action_type"`
	ActionDate *time.Time
}

type ReviewFilter struct {
	ReviewID    int
	BookID      int
	UserID      int
	CountOnPage int
	Page        int
	DateFrom    *time.Time
	DateTo      *time.Time
}
