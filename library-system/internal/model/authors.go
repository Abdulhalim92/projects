package model

type Author struct {
	Authors_id int `gorm:"primary_key"`
	Name       string
	Biography  string
}
