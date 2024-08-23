package model

type Review struct {
	reviews_id  int
	user_id     int
	book_id     int
	review_text string
	rating      float64
	review_data string
}
