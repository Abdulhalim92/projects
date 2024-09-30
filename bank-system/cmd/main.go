package main

import (
	"banking-system/bank-system/internal/config"
	"banking-system/bank-system/internal/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализируем базу данных
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	fmt.Println("Database initialized successfully:", db)

	// Запуск сервера
	log.Printf("Starting server on port %s...", cfg.Server.Port)
	if err := http.ListenAndServe(cfg.Server.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
