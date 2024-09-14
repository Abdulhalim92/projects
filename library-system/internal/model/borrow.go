package model

import "time"

type Borrow struct {
	BorrowID   int        `gorm:"column:borrow_id;primaryKey`
	UserID     int        `gorm:"column:user_id`
	BookID     int        `gorm:"column:book_id`
	BorrowDate *time.Time `gorm:"column:borrow_date`
	ReturnDate *time.Time `gorm:"column:return_date`
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
