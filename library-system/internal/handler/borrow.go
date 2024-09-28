package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) GetBorrows(c *gin.Context) {
	// Получение данных из БД через сервис
	borrows, err := h.service.GetBorrows()
	if err != nil {
		log.Printf("GetBorrows - h.service.GetBorrows error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetBorrows - borrows: %v", borrows)

	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) GetBorrowByUser(c *gin.Context) {
	// Получение ID пользователя из URL
	idStr := c.Param("id")
	if idStr == "" {
		log.Printf("GetBorrowByUser - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowByUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получение user_id из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		log.Printf("AddAuthor - user_id not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if id != userID.(int) {
		log.Printf("GetBorrowByUser - user_id doesn't match")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	borrows, err := h.service.GetBorrowsByUser(id)
	if err != nil {
		log.Printf("GetBorrowByUser - h.service.GetBorrowsByUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetBorrowByUser - borrows: %v", borrows)

	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) GetBorrowsByBook(c *gin.Context) {
	// Получение ID книги из URL
	idStr := c.Param("id")
	if idStr == "" {
		log.Printf("GetBorrowsByBook - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBorrowsByBook - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrows, err := h.service.GetBorrowsByBook(id)
	if err != nil {
		log.Printf("GetBorrowsByBook - h.service.GetBorrowsByBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetBorrowsByBook - borrows: %v", borrows)

	c.JSON(http.StatusOK, gin.H{"data": borrows})
}

func (h *Handler) CreateBorrow(c *gin.Context) {
	var borrow model.Borrow

	// Получение данных из тела запроса
	if err := c.BindJSON(&borrow); err != nil {
		log.Printf("CreateBorrow - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("CreateBorrow - data after binding: %v", borrow)

	// Получение ID пользователя из контекста
	userID, ok := c.Get("user_id")
	if !ok {
		log.Printf("CreateBorrow - user_id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// Проверка user_id в запросе
	if borrow.UserID != userID.(int) {
		log.Printf("CreateBorrow - user_id doesn't match")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	createdBorrow, err := h.service.CreateBorrow(&borrow)
	if err != nil {
		log.Printf("CreateBorrow - h.service.CreateBorrow error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("CreateBorrow - createdBorrow: %v", createdBorrow)

	c.JSON(http.StatusOK, gin.H{"data": createdBorrow})
}

func (h *Handler) ReturnBook(c *gin.Context) {
	// Получение ID книги из URL
	idStr := c.Param("id")
	if idStr == "" {
		log.Printf("ReturnBook - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ReturnBook - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получение user_id из контекста
	userID, ok := c.Get("user_id")
	if !ok {
		log.Printf("ReturnBook - user_id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if err := h.service.ReturnBook(userID.(int), id); err != nil {
		log.Printf("ReturnBook - h.service.ReturnBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ReturnBook - book with id %d returned by user %d", id, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Book returned"})
}
