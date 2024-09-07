package main

import (
	"fmt"
	"net/http"
	"projects/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		panic(err) // TODO
	}
	// Initialize repository, service, and handlers
	NewRepository := internal.NewRepository(db)
	NewService := internal.NewService(*NewRepository)
	mux := http.NewServeMux()
	NewHandler := internal.NewHandler(mux, NewService)
	NewHandler.InitRoutes()

	fmt.Printf("Server is starting... address: %v", ":8080\n")
	err = http.ListenAndServe("localhost:8080", NewHandler)
	if err != nil {
		panic(err)
	}
}

// connectToDB connects to the PostgreSQL database using the provided DSN.
// It returns a pointer to a gorm.DB object and an error if any occurred.
func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
