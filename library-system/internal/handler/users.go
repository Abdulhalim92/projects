package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *MyHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// получаем пользователей
	users, err := h.service.ListUsers()
	if err != nil {
		log.Printf("Failed to Get Users - service.ListUsers error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//преобразуем слайс юзеров в слайс байтов
	data, err := json.MarshalIndent(users, "   ", "")
	if err != nil {
		log.Printf("Failed to Get Users - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//объявляем тип контента и записываем респонс
	w.Header().Set("Content-Type", "json/application")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Get Users - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	//получаем айди из пути url и конвертируем в int
	id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		log.Printf("Failed to Get User By ID - strconv.Atoi error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//находим юзера с этим айди
	user, err := h.service.FindUser(id)
	if err != nil {
		log.Printf("Failed to Get User By ID - service.FindUser error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//convert it to json
	data, err := json.MarshalIndent(user, "   ", "")
	if err != nil {
		log.Printf("Failed to Get User By ID - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//write json to response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Get User By ID - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//read request body - we get json with a user structure
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//decode json
	var user *model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//update user
	user, err = h.service.EditUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err = json.MarshalIndent(user, "   ", "")
	if err != nil {
		log.Printf("Failed to Update User - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Update User - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *MyHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//decode json
	var user *model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//create user
	user, err = h.service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err = json.MarshalIndent(user, "   ", "")
	if err != nil {
		log.Printf("Failed to Add User - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Add User - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err = h.service.RemoveUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User with id %d was successfully deleted\n", id)

}
