package model

type Review struct {
	Reviews_id  int `gorm:"primary_key"`
	User_id     int
	Book_id     int
	Review_text string
	Rating      float64
	Review_data string
}
