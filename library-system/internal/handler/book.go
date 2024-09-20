package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.service.ListBooks()
	if err != nil {
		log.Printf("GetBooks - h.service.ListBooks error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBooks - data: %v", books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (h *Handler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetBookByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBookByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.FindBook(id)
	if err != nil {
		log.Printf("GetBookByID - h.service.FindBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBookByID - book: %v", book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (h *Handler) GetBooksByAuthor(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetBooksByAuthor - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBooksByAuthor - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books, err := h.service.GetBooksByAuthor(id)
	if err != nil {
		log.Printf("GetBooksByAuthor - h.service.GetBooksByAuthor error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetBooksByAuthor - books: %v", books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (h *Handler) AddBook(c *gin.Context) {
	var book model.Book

	if err := c.BindJSON(&book); err != nil {
		log.Printf("AddBook - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddBook - data after binding: %v", book)

	createBook, err := h.service.CreateBook(&book)
	if err != nil {
		log.Printf("AddBook - h.service.CreateBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddBook - created book: %v", createBook)
	c.JSON(http.StatusCreated, gin.H{"data": createBook})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var book model.Book

	if err := c.BindJSON(&book); err != nil {
		log.Printf("UpdateBook - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateBook - data after binding: %v", book)

	updatedBook, err := h.service.EditBook(&book)
	if err != nil {
		log.Printf("UpdateBook - h.service.EditBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateBook - updated book: %v", updatedBook)
	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteBook - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.RemoveBook(id); err != nil {
		log.Printf("DeleteBook - h.service.RemoveBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteBook - book with id %d deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
