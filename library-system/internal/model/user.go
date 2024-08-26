package model

type User struct {
	Users_id int `gorm:"primary_key"`
	Username string
	Password string
}
