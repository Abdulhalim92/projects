package model

import "time"

type User struct {
	UserID   int `gorm:"primaryKey"`
	Username string
	Password string
	//HasProfile bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Profile struct {
	UserID  int
	Email   string
	Address string
}
