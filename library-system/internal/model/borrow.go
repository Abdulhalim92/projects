package model

type Borrow struct {
	BorrowID   int    `gorm:"primaryKey:borrow_id"`
	UserID     int    `gorm:"column:user_id"`
	BookID     int    `gorm:"column:book_id"`
	BorrowDate string `gorm:"column:borrowdate"`
	ReturnDate string `gorm:"column:returndate"`
}
