package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"projects/internal/model"
	"projects/internal/service"
	"strconv"
)

type UserHandler struct {
	mux     *http.ServeMux
	service *service.UserService
}

func NewUserHandler(mux *http.ServeMux, s *service.UserService) *UserHandler {
	return &UserHandler{mux: mux, service: s}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *UserHandler) InitRoutes() {
	h.mux.HandleFunc("/users", h.GetUsers)
	h.mux.HandleFunc("/users/{user_id}", h.GetUserByID)
	h.mux.HandleFunc("/users/update", h.UpdateUser)
	h.mux.HandleFunc("/users/add", h.AddUser)
	h.mux.HandleFunc("/users/{user_id}/delete", h.DeleteUser)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// получаем пользователей
	users, err := h.service.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//преобразуем слайс юзеров в слайс байтов
	data, err := json.MarshalIndent(users, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//объявляем тип контента и записываем респонс
	w.Header().Set("Content-Type", "json/application")
	w.Write(data)

}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	//получаем айди из пути url и конвертируем в int
	id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//находим юзера с этим айди
	user, err := h.service.FindUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//convert it to json
	data, err := json.MarshalIndent(user, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//write json to response
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//read request body - we get json with a user structure
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//decode json
	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//update user
	h.service.EditUser(&user)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//decode json
	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//create user
	h.service.CreateUser(&user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.service.RemoveUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
