package models
import "time"
// Account - модель для банковских счетов
type Account struct {
    AccountID     int     `json:"account_id,omitempty" gorm:"primaryKey"` // Уникальный идентификатор счета
    AccountNumber string  `json:"account_number,omitempty"`               // Номер банковского счета
    AccountType   string  `json:"account_type,omitempty"`                 // Тип банковского счета (например, текущий или сберегательный)
    Balance       float64 `json:"balance,omitempty"`                      // Текущий баланс счета
    OwnerID       int     `json:"owner_id,omitempty"`                     // Внешний ключ на владельца (связь с пользователем)
	CreatedAt     time.Time `json:"created_at,omitempty"`                 // Время создания счета
    UpdatedAt     time.Time `json:"updated_at,omitempty"`                 // Время последнего обновления
}
