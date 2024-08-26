package model

type Profile struct {
	User_id int `gorm:"primary_key"`
	Email   string
	Address string
}
