package model

type User struct {
	UserId   int `gorm:"column:id;primaryKey"`
	Username string
	Password string
}

type Profile struct {
	UserId  int
	Email   string
	Address string
}
