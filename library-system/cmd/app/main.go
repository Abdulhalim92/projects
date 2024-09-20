package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"projects/internal/handler"
	"projects/internal/repository"
	"projects/internal/service"
)

func main() {
	fmt.Println("Library System")

	// Подключение к базе данных
	db, err := ConnectToDb()
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
func ConnectToDb() (*gorm.DB, error) {
	dsn := "host=localhost user=humo password=humo dbname=Humo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return nil, err
	}
	return db, nil
}
