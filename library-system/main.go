package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"projects/internal/model"
)

type User struct {
	ID       int `gorm:"column:user_id;primaryKey"`
	Username string
	Password string
}

func main() {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Миграция схемы
	//db.AutoMigrate(&User{})
	var u []User
	err = db.Find(&u).Error
	if err != nil {
		panic(err)
	}

	var us User
	err = db.Last(&us).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(us)

	for _, user := range u {
		fmt.Printf("id: %d username: %s password: %s\n", user.ID, user.Username, user.Password)
	}

	db.First(&us).Where("user_id = ?", 1)
	db.Raw("SELECT * FROM users WHERE user_id = 1").Scan(&us)

	//	SELECT b.title FROM books b
	//JOIN borrow br USING (book_id)
	//WHERE br.return_date > '2022-01-30';
	db.Model(&model.Book{}).Select("title").
		Joins("JOIN borrow br USING (book_id)").
		Where("br.return_date > ?", "2022-01-30").
		Scan(&u)
}
