package model

type Author struct {
	AuthorID  int `gorm:"column:author_id;primaryKey"`
	Name      string
	Biography string
}
