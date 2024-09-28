package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("SignUp - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("SignUp - data after unmarshalling: %v", user)

	createUser, err := h.service.CreateUser(&user)
	if err != nil {
		log.Printf("SignUp - h.service.CreateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("SignUp - response to client: %v", string(createUser))

	c.JSON(http.StatusOK, gin.H{"data": createUser})
}

func (h *Handler) SignIn(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		log.Printf("SignIn - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("SignIn - data after binding: %v", user)

	token, err := h.service.SignIn(&user)
	if err != nil {
		log.Printf("SignIn - h.service.SignIn error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("SignIn - response to client: %v", string(token))

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		log.Printf("GetUsers - h.service.ListUsers error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	// Получение ID пользователя из URL
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetUserByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetUserByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.FindUser(id)
	if err != nil {
		log.Printf("GetUserByID - h.service.GetUserByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		log.Printf("UpdateUser - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser, err := h.service.EditUser(&user)
	if err != nil {
		log.Printf("UpdateUser - h.service.UpdateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updateUser})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("DeleteUser - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteUser(id); err != nil {
		log.Printf("DeleteUser - h.service.DeleteUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
