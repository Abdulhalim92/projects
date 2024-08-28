package model

type Author struct {
	Authorid  int `gorm:"primaryKey"`
	Name      string
	Biography string
}
