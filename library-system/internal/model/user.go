package model

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"
)

// User структура для пользователя.
type User struct {
	UserID   int    `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	//HasProfile bool
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ValidateUser функция, которая проверяет данные пользователя на соответствие заданным критериям.
func (u *User) ValidateUser() error {
	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	if err := u.validatePassword(); err != nil {
		return err
	}

	return nil
}

// validatePassword проверяет сложность пароля.
func (u *User) validatePassword() error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(u.Password) >= 8 {
		hasMinLen = true
	}

	for _, char := range u.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

// Profile структура для профиля пользователя.
type Profile struct {
	UserID  int    `json:"user_id" gorm:"primaryKey"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

// ValidateProfile функция, которая проверяет данные профиля пользователя на соответствие заданным критериям.
func (p *Profile) ValidateProfile() error {
	if !p.isValidEmail() {
		return errors.New("invalid email format")
	}

	if len(p.Address) == 0 {
		return errors.New("address cannot be empty")
	}

	return nil
}

// isValidEmail проверяет корректность email с использованием регулярного выражения.
func (p *Profile) isValidEmail() bool {
	// Пример простого регулярного выражения для проверки email
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(p.Email)
}
