package model

import "time"

type Book struct {
	BookId   int `gorm:"column:id;primaryKey"`
	Title    string
	AuthorId int `gorm:"foreignKey"`
}

type Review struct {
	ReviewId   int
	UserId     int
	BookId     int
	ReviewText string
	Rating     float64
	ReviewDate time.Time
}

type Borrow struct {
	BorrowId   int
	UserId     int
	BookId     int
	BorrowDate time.Time
	ReturnDate time.Time
}
