package user

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
)

type UserHandler struct {
	mux         *http.ServeMux
	userService *Service
}

func NewUserHandler(mux *http.ServeMux, s *Service) *UserHandler {
	return &UserHandler{
		mux:         mux,
		userService: s,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for users to the given ServeMux.
//
// It adds the following routes:
//
// - GET /users: GetUsers
// - GET /users/:id: GetUserByID
// - POST /users/add: AddUser
// - DELETE /users/delete/:id: DeleteUser
func (h *UserHandler) InitRoutes() {
	h.mux.HandleFunc("/users", h.GetUsers)
	h.mux.HandleFunc("/users/:id", h.GetUserByID)
	h.mux.HandleFunc("/users/add", h.AddUser)
	h.mux.HandleFunc("/users/delete/:id", h.DeleteUser)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling users: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := h.userService.FindUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonUser, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling user: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
} // TODO: implement

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user = &model.User{
		Username: username,
		Password: password,
	}

	user, err = h.userService.ur.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonUser, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)

} // TODO: implement

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.userService.RemoveUser(userID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("User with id:%d removed successfully", userID)
	w.Write([]byte(response))
} // TODO: implement
