package model

type User struct {
	Userid   int `gorm:"primaryKey"`
	Username string
	Password string
}
