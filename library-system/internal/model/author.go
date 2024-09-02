package model

type Author struct {
	AuthorId  int `gorm:"column:id;primaryKey"`
	Name      string
	Biography string
	Address   string
}
