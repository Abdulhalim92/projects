package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"projects/internal/repository"
	"projects/internal/service"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

}

// Подключение к базе данных
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=e-donish port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
