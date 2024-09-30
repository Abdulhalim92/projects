package models
import "time"

// Transaction - модель для операций с банковским счетом
type Transaction struct {
    TransactionID   int     `json:"transaction_id,omitempty" gorm:"primaryKey"` // Уникальный идентификатор транзакции
    AccountID       int     `json:"account_id,omitempty"`                      // Внешний ключ на банковский счет
    TransactionType string  `json:"transaction_type,omitempty"`                // Тип транзакции: депозит, снятие, перевод
    Amount          float64 `json:"amount,omitempty"`                          // Сумма транзакции
    TransactionDate string  `json:"transaction_date,omitempty"`                // Дата транзакции
    Description     string  `json:"description,omitempty"`                     // Описание транзакции
	CreatedAt       time.Time `json:"created_at,omitempty"`                   // Время создания записи о транзакции
    UpdatedAt       time.Time `json:"updated_at,omitempty"`                  // Время последнего обновления
}
