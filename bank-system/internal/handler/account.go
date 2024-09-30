package handlers

import (
	"banking-system/bank-system/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AccountHandler - структура для обработчиков запросов, связанных с аккаунтами
type AccountHandler struct {
	Service *service.AccountService
}

// CreateAccountHandler - обработчик для создания нового счета
func (h *AccountHandler) CreateAccountHandler(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateAccount(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// GetBalanceHandler - обработчик для получения баланса счета
func (h *AccountHandler) GetBalanceHandler(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	balance, err := h.Service.GetBalance(accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// DepositHandler - обработчик для пополнения счета
func (h *AccountHandler) DepositHandler(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	var requestBody struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Deposit(accountID, requestBody.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deposit successful"})
}

// WithdrawHandler - обработчик для снятия средств
func (h *AccountHandler) WithdrawHandler(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("accountID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	var requestBody struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Withdraw(accountID, requestBody.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
}
