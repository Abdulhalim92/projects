package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) GetAuthors(c *gin.Context) {
	authors, err := h.service.GetAuthors()
	if err != nil {
		log.Printf("GetAuthors - h.service.ListAuthors error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetAuthors - authors: %v", authors)
	c.JSON(http.StatusOK, gin.H{"data": authors})
}

func (h *Handler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetAuthorByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetAuthorByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		log.Printf("GetAuthorByID - h.service.GetAuthorByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetAuthorByID - author: %v", author)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

func (h *Handler) AddAuthor(c *gin.Context) {
	var author model.Author

	if err := c.BindJSON(&author); err != nil {
		log.Printf("AddAuthor - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddAuthor - data after binding: %v", author)

	createAuthor, err := h.service.CreateAuthor(&author)
	if err != nil {
		log.Printf("AddAuthor - h.service.CreateAuthor error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddAuthor - created author: %v", createAuthor)
	c.JSON(http.StatusOK, gin.H{"data": createAuthor})
}

func (h *Handler) UpdateAuthor(c *gin.Context) {
	var author model.Author

	if err := c.BindJSON(&author); err != nil {
		log.Printf("UpdateAuthor - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateAuthor - data after binding: %v", author)

	updateAuthor, err := h.service.EditAuthor(&author)
	if err != nil {
		log.Printf("UpdateAuthor - h.service.EditAuthor error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateAuthor - updated author: %v", updateAuthor)
	c.JSON(http.StatusOK, gin.H{"data": updateAuthor})
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteAuthor - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteAuthor(id); err != nil {
		log.Printf("DeleteAuthor - h.service.DeleteAuthor error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteAuthor - author with id %d deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
