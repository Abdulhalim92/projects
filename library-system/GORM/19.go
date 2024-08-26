package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func ShowUsersWithMoretakingbooksthanone(db *gorm.DB) {
	var x []model.User
	db.Raw(" SELECT u.username, COUNT(br.book_id) FROM users u JOIN borrow br ON u.users_id = br.user_id GROUP BY u.username HAVING COUNT(br.book_id) > 1;").Scan(&x)
	fmt.Println(x)
}
