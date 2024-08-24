package model

type User struct {
	ID       int `gorm:"column:user_id;primaryKey"`
	Username string
	Password string
}

type Profiles struct {
	UserID  int
	Email   string
	Address string
}
