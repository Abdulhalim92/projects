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
	Service *service.Service
}

func (u *UserHandler) InitUsers() {
	u.mux.HandleFunc("/users", u.AddUser)
	u.mux.HandleFunc("/users/{id}", u.GetUserById)
	u.mux.HandleFunc("/users/delete/{id}", u.DeleteUser)
	u.mux.HandleFunc("/users/update", u.DeleteUser)
}
func CreateUserHandler(mux *http.ServeMux, Service *service.Service) *UserHandler {
	return &UserHandler{
		mux:     mux,
		Service: Service,
	}
}
func (u *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = u.Service.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Service.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, err := json.MarshalIndent(users, " ", "    ")
	if err != nil {
		http.Error(w, err.Error(), 0777)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user, err := u.Service.ListUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, err := json.MarshalIndent(user, " ", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = u.Service.EditUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Upgraded"))
}
func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = u.Service.RemoveUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Delete"))
}
