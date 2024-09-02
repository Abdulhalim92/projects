package model

type Book struct {
	BookID   int    `gorm:"primaryKey:book_id"`
	Title    string `gorm:"column:title"`
	AuthorID int    `gorm:"column:author_id"`
}
