package model

import "time"

type Review struct {
	ReviewID   int        `gorm:"column:review_id;primaryKey"`
	UserID     int        `gorm:"column:user_id"`
	BookID     int        `gorm:"column:book_id"`
	ReviewText string     `gorm:"column:review_text"`
	Rating     float32    `gorm:"column:rating"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time
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
