package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *Handler) ListProfiles(c *gin.Context) {
	// Получение всех профилей
	profiles, err := h.service.ListProfiles()
	if err != nil {
		log.Printf("ListProfiles - h.service.ListProfiles error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("ListProfiles - profiles: %v", profiles)

	c.JSON(http.StatusOK, gin.H{"data": profiles})
}

func (h *Handler) CreateProfile(c *gin.Context) {
	var profile model.Profile

	// Получение данных из тела запроса
	if err := c.BindJSON(&profile); err != nil {
		log.Printf("CreateProfile - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("CreateProfile - data after binding: %v", profile)

	createProfile, err := h.service.CreateProfile(&profile)
	if err != nil {
		log.Printf("CreateProfile - h.service.CreateProfile error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("CreateProfile - created profile: %v", createProfile)

	c.JSON(http.StatusOK, gin.H{"data": createProfile})
}

func (h *Handler) GetProfileByID(c *gin.Context) {
	idStr := c.Param("id")

	// Получение ID профиля из URL
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetProfileByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.GetProfileByID(id)
	if err != nil {
		log.Printf("GetProfileByID - h.service.GetProfileByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("GetProfileByID - profile: %v", profile)

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var profile model.Profile

	// Получение данных из тела запроса
	if err := c.BindJSON(&profile); err != nil {
		log.Printf("UpdateProfile - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("UpdateProfile - data after binding: %v", profile)

	updatedProfile, err := h.service.EditProfile(&profile)
	if err != nil {
		log.Printf("UpdateProfile - h.service.EditProfile error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("UpdateProfile - updated profile: %v", updatedProfile)

	c.JSON(http.StatusOK, gin.H{"data": updatedProfile})
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	// Получение ID профиля из URL
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteProfile - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteProfile(id); err != nil {
		log.Printf("DeleteProfile - h.service.DeleteProfile error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("DeleteProfile - profile with id %d deleted", id)

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
