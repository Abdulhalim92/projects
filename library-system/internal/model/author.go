package model

import "time"

type Author struct {
	AuthorID  int        `json:"author_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	Biography string     `json:"biography,omitempty"`
	Address   string     `json:"address,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
