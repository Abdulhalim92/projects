package model

import "time"

type Book struct {
	BookId   int    `json:"bookd_id" gorm:"primaryKey"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

type Reviews struct {
	ReviewID   int
	UserID     int
	BookID     int
	ReviewText string
	Rating     float64
	ReviewDate time.Time
}

type Borrow struct {
	BorrowID   int
	UserID     int
	BookID     int
	BorrowDate time.Time
	ReturnDate time.Time
}
