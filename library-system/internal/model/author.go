package model

type Author struct {
	AuthorID  int    `json:"author_id,omitempty" gorm:"primaryKey"`
	Name      string `json:"name,omitempty"`
	Biography string `json:"biography,omitempty"`
	Address   string `json:"address,omitempty"`
}
