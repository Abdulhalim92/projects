package model

type Review struct {
	Reviewid   int `gorm:"primaryKey"`
	Userid     int
	Bookid     int
	ReviewText string
	Rating     float64
	ReviewData string
}
