package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"projects/internal/handler"
	"projects/internal/model"
	"projects/internal/repository"
	"projects/internal/service"
	"time"
)

func OperateThroughCL() {
	for {
		var command string

		fmt.Scan(&command)
		if command == "stop" {
			break
		}
	}
}

func main() {
	//cd library-system/cmd/app
	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		fmt.Errorf("Failed to connect to Database: %v\n", err)
	}

	mux := http.NewServeMux()

	// Инициализация книг
	Repository := repository.NewRepository(db)
	Service := service.NewService(*Repository)

	handler := handler.NewMyHandler(mux, Service)
	handler.InitRoutes()

	fmt.Printf("Server is starting...address: %s", ":8080\n")

	t := time.Now()
	borrow := &model.Borrow{BookID: 5, UserID: 3, BorrowDate: &t}
	Repository.AddBorrow(borrow)

	user := &model.User{Username: "mike", Password: "yoo"}
	Repository.AddUser(user)
	//http.ListenAndServe("localhost:8080", handler)

}

func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=muqaddas password=password dbname=library_db port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %v\n", err)
	}
	return db, nil
}
