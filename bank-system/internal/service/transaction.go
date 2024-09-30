package services

import (
    "banking-system/bank-system/internal/model"
    "gorm.io/gorm"
)

// TransactionService - структура для сервиса транзакций
type TransactionService struct {
    DB *gorm.DB
}

// CreateTransaction - функция для создания новой транзакции
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
    // Создание новой транзакции
    return s.DB.Create(transaction).Error
}

// GetTransactionHistory - функция для получения истории транзакций по счету
func (s *TransactionService) GetTransactionHistory(accountID int) ([]models.Transaction, error) {
    var transactions []models.Transaction
    if err := s.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}
