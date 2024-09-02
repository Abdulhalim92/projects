package model

type User struct {
	UserID   int    `gorm:"primaryKey:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
