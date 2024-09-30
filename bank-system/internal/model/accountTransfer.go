package models
import "time"
// AccountTransfer - модель для переводов между счетами
type AccountTransfer struct {
    TransferID    int     `json:"transfer_id,omitempty" gorm:"primaryKey"` // Уникальный идентификатор перевода
    FromAccountID int     `json:"from_account_id,omitempty"`               // Счет отправителя
    ToAccountID   int     `json:"to_account_id,omitempty"`                 // Счет получателя
    TransferAmount float64 `json:"transfer_amount,omitempty"`              // Сумма перевода
	TransferDate   time.Time `json:"transfer_date,omitempty"`              // Дата и время перевода
    CreatedAt      time.Time `json:"created_at,omitempty"`                 // Время создания записи о переводе
    UpdatedAt      time.Time `json:"updated_at,omitempty"`                 // Время последнего обновления
}
