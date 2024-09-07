package model

import "time"

type Book struct {
	BookID   int    `json:"book_id" gorm:"primaryKey"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

type Reviews struct {
	ReviewID   int       `json:"review_id" gorm:"primaryKey"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	ReviewText string    `json:"review_text"`
	Rating     float64   `json:"rating"`
	ReviewDate time.Time `json:"review_date"`
}

type Borrow struct {
	BorrowID   int       `json:"borrow_id" gorm:"primaryKey"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

type ReviewFilter struct {
	ReviewID int
	BookID   int
	UserID   int
}
