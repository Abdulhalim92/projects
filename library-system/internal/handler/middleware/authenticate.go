package middleware

import (
	"context"
	"log"
	"net/http"
	"projects/internal/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка
		authToken := r.Header.Get("Authorization")

		// Проверяем наличие токена
		if authToken == "" {
			log.Printf("Authenticate - authToken is empty")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Получаем ID пользователя из JWT
		userID, err := utils.ValidateJWT(authToken)
		if err != nil {
			log.Printf("Authenticate - utils.ValidateJWT error: %v", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		// Добавляем ID пользователя в контекст
		ctx := context.WithValue(r.Context(), "user_id", userID)

		// Добавляем контекст в запрос
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
