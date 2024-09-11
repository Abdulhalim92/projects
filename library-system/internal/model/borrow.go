package model

import "time"

type Borrow struct {
	BorrowID   int
	UserID     int
	BookID     int
	BorrowDate *time.Time
	ReturnDate *time.Time
}
