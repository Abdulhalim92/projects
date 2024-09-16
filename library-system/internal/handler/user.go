package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
	"strings"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		log.Printf("GetUsers - h.service.ListUsers error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Printf("GetUsers - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	if idStr == "" {
		log.Printf("GetUserByID - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetUserByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.FindUser(id)
	if err != nil {
		log.Printf("GetUserByID - h.service.GetUserByID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Printf("GetUserByID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

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

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("UpdateUser - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateUser - incoming request: %v", string(data))

	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Printf("UpdateUser - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateUser - data after unmarshalling: %v", user)

	updateUser, err := h.service.EditUser(&user)
	if err != nil {
		log.Printf("UpdateUser - h.service.UpdateUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateUser - updated user: %v", updateUser)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(updateUser, "", "    ")
	if err != nil {
		log.Printf("UpdateUser - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateUser - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/delete/")

	if idStr == "" {
		log.Printf("DeleteUser - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteUser - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.DeleteUser(id)
	if err != nil {
		log.Printf("DeleteUser - h.service.DeleteUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
