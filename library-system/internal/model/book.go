package model

type Book struct {
	Books_id  int `gorm:"primary_key"`
	Title     string
	Author_id int
}
