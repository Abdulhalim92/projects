package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	//
	mux := http.NewServeMux()
	// Инициализация обработчика
	newHandler := handler.NewHandler(mux, newService)
	newHandler.InitRoutes()

	fmt.Printf("Server is starting... address: %v", ":8080\n")
	err = http.ListenAndServe("localhost:8080", newHandler)
	if err != nil {
		panic(err)
	}
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
