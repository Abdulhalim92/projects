package model

import "time"

type Review struct {
	ReviewID   int       `gorm:"column:review_id;primaryKey"`
	UserID     int       `gorm:"column:user_id"`
	BookID     int       `gorm:"column:book_id"`
	ReviewText string    `gorm:"column:review_text"`
	Rating     int       `gorm:"column:rating"`
	ReviewDate time.Time `gorm:"column:review_date"`
}

type ReviewFilter struct {
	ReviewID    int
	BookID      int
	UserID      int
	CountOnPage int
	Page        int
	DateFrom    time.Time
	DateTo      time.Time
}
