package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) GetReviews(c *gin.Context) {
	// Получение всех отзывов
	reviews, err := h.service.ListReviews()
	if err != nil {
		log.Printf("GetReviews - h.service.ListReviews error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviews - reviews: %v", reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (h *Handler) AddReview(c *gin.Context) {
	var review model.Reviews

	// Получение данных из тела запроса
	if err := c.BindJSON(&review); err != nil {
		log.Printf("AddReview - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("AddReview - data after binding: %v", review)

	createReview, err := h.service.CreateReview(&review)
	if err != nil {
		log.Printf("AddReview - h.service.CreateReview error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("AddReview - created review: %v", createReview)

	c.JSON(http.StatusOK, gin.H{"data": createReview})
}

func (h *Handler) GetReviewByID(c *gin.Context) {
	// Получение ID отзыва из URL
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReviewByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.GetReviewByID(id)
	if err != nil {
		log.Printf("GetReviewByID - h.service.GetReviewByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviewByID - review: %v", review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

func (h *Handler) GetReviewsByBook(c *gin.Context) {
	// Получение ID книги из URL
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReviewsByBook - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviews, err := h.service.GetReviewsByBook(id)
	if err != nil {
		log.Printf("GetReviewsByBook - h.service.GetReviewsByBook error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviewsByBook - reviews: %v", reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (h *Handler) GetReviewsByUser(c *gin.Context) {
	// Получение ID пользователя из URL
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetReviewsByUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviews, err := h.service.GetReviewsByUser(id)
	if err != nil {
		log.Printf("GetReviewsByUser - h.service.GetReviewsByUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviewsByUser - reviews: %v", reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (h *Handler) GetReviewsByFilter(c *gin.Context) {
	var filter model.ReviewFilter

	// Получение данных из тела запроса
	if err := c.BindJSON(&filter); err != nil {
		log.Printf("GetReviewsByFilter - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviewsByFilter - data after unmarshalling: %v", filter)

	reviews, err := h.service.GetReviewsByFilter(filter)
	if err != nil {
		log.Printf("GetReviewsByFilter - h.service.GetReviewsByFilter error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetReviewsByFilter - reviews: %v", reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (h *Handler) UpdateReview(c *gin.Context) {
	var review model.Reviews

	// Получение данных из тела запроса
	if err := c.BindJSON(&review); err != nil {
		log.Printf("UpdateReview - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("UpdateReview - data after binding: %v", review)

	updatedReview, err := h.service.EditReview(&review)
	if err != nil {
		log.Printf("UpdateReview - h.service.EditReview error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("UpdateReview - updated review: %v", updatedReview)

	c.JSON(http.StatusOK, gin.H{"data": updatedReview})
}

func (h *Handler) DeleteReview(c *gin.Context) {
	// Получение ID отзыва из URL
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteReview - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteReview(id); err != nil {
		log.Printf("DeleteReview - h.service.DeleteReview error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("DeleteReview - review with id %d deleted", id)

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}
