package model

type Review struct {
	ReviewID   int `gorm:"primaryKey"`
	UserID     int
	BookID     int
	ReviewText string
	Rating     float64
	ReviewData string
}
