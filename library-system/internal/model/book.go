package model

type Book struct {
	Bookid   int `gorm:"primaryKey"`
	Title    string
	Authorid int
}
