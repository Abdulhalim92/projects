package main

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func ShowUsersWithMoretakingbooksthanone(db *gorm.DB) {
	var x []model.User
	db.Raw(" SELECT u.username, COUNT(br.bookId) FROM users u JOIN borrow br ON u.userId = br.userId GROUP BY u.username HAVING COUNT(br.bookId) > 1;").Scan(&x)
	fmt.Println(x)
}
