package model

import "time"

type Book struct {
	BookId   int `gorm:"column:id;primaryKey"`
	Title    string
	AuthorId int `gorm:"foreignKey"`
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
