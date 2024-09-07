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
	db, err := connectToDB()
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

// connectToDB connects to the PostgresSQL database using the provided DSN.
//
// It returns a pointer to a gorm.DB object and an error if any occurred.
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
