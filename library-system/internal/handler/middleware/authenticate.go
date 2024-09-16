package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		authToken, ok := c.Get("Authorization")
		if !ok {
			log.Printf("Authenticate - authToken is empty")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		strToken := authToken.(string)

		// Получаем ID пользователя из JWT
		userID, err := utils.ValidateJWT(strToken)
		if err != nil {
			log.Printf("Authenticate - utils.ValidateJWT error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Добавляем ID пользователя в контекст
		c.Set("user_id", userID)

		c.Next()
	}
}
