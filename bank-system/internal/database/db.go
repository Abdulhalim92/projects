package database

import (
	"banking-system/bank-system/internal/config"
	"banking-system/bank-system/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// InitDB - функция для инициализации базы данных с использованием конфигурации
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	// Автоматическая миграция
	err = db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.User{}, &models.AccountTransfer{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return nil, err
	}

	log.Println("Database successfully connected and migrated")
	return db, nil
}
