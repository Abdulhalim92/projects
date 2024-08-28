package model

type Borrow struct {
	Borrowid   int `gorm:"primaryKey"`
	Userid     int
	Bookid     int
	BorrowDate string
	ReturnDate string
}
