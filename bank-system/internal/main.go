package main

import (
	"banking-system/bank-system/internal/handler"
	"banking-system/bank-system/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Инициализация базы данных
	dsn := "host=localhost user=username password=password dbname=bank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автоматическая миграция (создание таблиц)
	db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.User{}, &models.AccountTransfer{})

	// Инициализация сервисов
	accountService := &service.AccountService{DB: db}
	transactionService := &service.TransactionService{DB: db}

	// Инициализация обработчиков
	accountHandler := &handlers.AccountHandler{Service: accountService}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}

	// Создание роутера Gin
	r := gin.Default()

	// Эндпоинты для аккаунтов
	r.POST("/accounts", accountHandler.CreateAccountHandler)
	r.GET("/accounts/:accountID/balance", accountHandler.GetBalanceHandler)
	r.POST("/accounts/:accountID/deposit", accountHandler.DepositHandler)
	r.POST("/accounts/:accountID/withdraw", accountHandler.WithdrawHandler)

	// Эндпоинты для транзакций
	r.GET("/accounts/:accountID/transactions", transactionHandler.GetTransactionHistoryHandler)
	r.POST("/transactions", transactionHandler.CreateTransactionHandler)

	// Запуск сервера
	r.Run(":8080")
}
