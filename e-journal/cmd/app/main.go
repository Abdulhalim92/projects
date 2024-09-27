package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"projects/internal/model"
	"projects/internal/repository"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	newRepository := repository.NewRepository(db)

	newUser := &model.User{
		Username: "shohin",
		Password: "admin",
		RoleID:   1,
	}

	userID, err := newRepository.CreateUser(newUser)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	log.Printf("User created with id: %v", userID)

	users, err := newRepository.GetUsers()
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	log.Printf("Users: %v", users)
}

// Подключение к базе данных
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=e_donish port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
