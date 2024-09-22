package model

import "time"

type User struct {
	UserID    int        `json:"user_id,omitempty" gorm:"primaryKey"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	RoleID    int        `json:"role_id,omitempty"` // Связь с таблицей ролей
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Role struct {
	RoleID    int        `json:"role_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"` // Название роли, например "admin", "teacher", "parent"
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Parent struct {
	ParentID  int        `json:"parent_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	UserID    int        `json:"user_id,omitempty"`    // Связь с таблицей пользователей
	StudentID int        `json:"student_id,omitempty"` // Связь с учеником
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Teacher struct {
	TeacherID int        `json:"teacher_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	UserID    int        `json:"user_id,omitempty"` // Связь с таблицей пользователей
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Admin struct {
	AdminID   int        `json:"admin_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	UserID    int        `json:"user_id,omitempty"` // Связь с таблицей пользователей
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
