package model

import "time"

type Book struct {
	BookID   int `gorm:"column:book_id;primaryKey"`
	Title    string
	AuthorID int `gorm:"column:author_id"`
}

type borrow struct {
	ID         int       `gorm:"column:id;primaryKey"`
	UserID     int       `gorm:"column:user_id"`
	BookID     int       `gorm:"column:book_id"`
	BorrowDate time.Time `gorm:"column:borrow_date"`
	ReturnDate time.Time `gorm:"column:return_date"`
}
