package model

type Author struct {
	AuthorID  int `gorm:"primaryKey"`
	Name      string
	Biography string
	Address   string
}
