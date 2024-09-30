package services

import (
    "errors"
    "banking-system/bank-system/internal/model" 
    "gorm.io/gorm"
)

// AccountService - структура для сервиса аккаунтов
type AccountService struct {
    DB *gorm.DB
}

// CreateAccount - функция для создания нового банковского счета
func (s *AccountService) CreateAccount(account *models.Account) error {
    // Проверка на уникальность номера счета
    var existingAccount models.Account
    if err := s.DB.Where("account_number = ?", account.AccountNumber).First(&existingAccount).Error; err == nil {
        return errors.New("account number already exists")
    }
    
    // Создание нового счета
    return s.DB.Create(account).Error
}

// GetBalance - функция для получения баланса счета
func (s *AccountService) GetBalance(accountID int) (float64, error) {
    var account models.Account
    if err := s.DB.First(&account, accountID).Error; err != nil {
        return 0, err
    }
    return account.Balance, nil
}

// Deposit - функция для пополнения счета
func (s *AccountService) Deposit(accountID int, amount float64) error {
    var account models.Account
    if err := s.DB.First(&account, accountID).Error; err != nil {
        return err
    }

    account.Balance += amount
    return s.DB.Save(&account).Error
}

// Withdraw - функция для снятия средств со счета
func (s *AccountService) Withdraw(accountID int, amount float64) error {
    var account models.Account
    if err := s.DB.First(&account, accountID).Error; err != nil {
        return err
    }

    if account.Balance < amount {
        return errors.New("insufficient funds")
    }

    account.Balance -= amount
    return s.DB.Save(&account).Error
}
