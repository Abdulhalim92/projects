package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"projects/internal/handler"
	"projects/internal/repository"
	"projects/internal/service"
)

func main() {
	fmt.Println("Library System")

	// Подключение к базе данных
	db, err := connectToDB()
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	// Инициализация репозитория
	newRepository := repository.NewRepository(db)
	// Инициализация сервиса
	newService := service.NewService(*newRepository)
	// Инициализация роутера
	router := gin.Default()
	// Инициализация обработчика
	newHandler := handler.NewHandler(router, newService)
	newHandler.InitRoutes()

	log.Fatalf("Failed to start server: %v", router.Run(":8080"))
}

// Подключение к базе данных
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
