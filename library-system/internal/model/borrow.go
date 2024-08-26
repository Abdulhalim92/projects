package model

type Borrow struct {
	Borrow_id   int `gorm:"primary_key"`
	User_id     int
	Book_id     int
	Borrow_date string
	Return_date string
}
