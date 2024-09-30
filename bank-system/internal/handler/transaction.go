package handlers

import (
	"banking-system/bank-system/internal/model"
	"banking-system/bank-system/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TransactionHandler - структура для обработчиков запросов, связанных с транзакциями
type TransactionHandler struct {
	Service *service.TransactionService
}

// GetTransactionHistoryHandler - обработчик для получения истории транзакций
func (h *TransactionHandler) GetTransactionHistoryHandler(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	transactions, err := h.Service.GetTransactionHistory(accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// CreateTransactionHandler - обработчик для создания новой транзакции
func (h *TransactionHandler) CreateTransactionHandler(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
