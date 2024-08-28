package model

type User struct {
	UserID   int `gorm:"column:id;primaryKey"`
	Username string
	Password string
}

type Profiles struct {
	UserID  int
	Email   string
	Address string
}
