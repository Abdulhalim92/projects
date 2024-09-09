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
	db, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	b := repository.NewBookRepository(db)
	BookService := service.NewService(b)
	BookHandler := handler.NewBookHandler(mux, BookService)
	BookHandler.InitRoutes()
	fmt.Printf("server is starring on port: %v", ":8080/n")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
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
