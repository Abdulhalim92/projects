package models
import "time"
// User - модель для пользователей (владельцев счетов)
type User struct {
    UserID    int    `json:"user_id,omitempty" gorm:"primaryKey"` // Уникальный идентификатор пользователя
    Username  string `json:"username,omitempty"`                  // Имя пользователя
    Email     string `json:"email,omitempty"`                     // Электронная почта пользователя
    Password  string `json:"password,omitempty"`                  // Пароль пользователя (зашифрованный)
	CreatedAt time.Time `json:"created_at,omitempty"`             // Время создания учетной записи
    UpdatedAt time.Time `json:"updated_at,omitempty"`             // Время последнего обновления
}
