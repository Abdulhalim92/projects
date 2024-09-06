package model

type User struct {
	UserID   int `gorm:"column:user_id;primaryKey"`
	Username string
	Password string
	//HasProfile bool
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

type Profile struct {
	UserID  int `gorm:"column:user_id;primaryKey"`
	Email   string
	Address string
}
