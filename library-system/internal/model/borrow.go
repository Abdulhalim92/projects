package model

import "time"

type Borrow struct {
	BorrowID   int `gorm:"column:borrow_id;primaryKey"`
	UserID     int
	BookID     int
	BorrowDate *time.Time `gorm:"column:borrow_date"`
	ReturnDate *time.Time
}

type BorrowFilter struct {
	BorrowID    int
	UserID      int
	BookID      int
	DateFrom    time.Time
	DateTo      time.Time
	WasReturned bool
	CountOnPage int
	Page        int
}
