package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
	"strings"
)

type Handler struct {
	mux     *http.ServeMux
	service *Service
}

func NewHandler(mux *http.ServeMux, s *Service) *Handler {
	return &Handler{
		mux:     mux,
		service: s,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) InitRoutes() {
	{
		h.mux.HandleFunc("/books", h.GetBooks)
		h.mux.HandleFunc("/books/{id}", h.GetBookByID)
		h.mux.HandleFunc("/books/add", h.AddBook)
		h.mux.HandleFunc("/books/delete/{id}", h.DeleteBook)
		h.mux.HandleFunc("/books/update/{id}", h.UpdateBook)
		h.mux.HandleFunc("/books/author/{id}", h.GetBooksByAuthor)
	}
	{
		h.mux.HandleFunc("/users", h.GetUsers)
		h.mux.HandleFunc("/users/{id}", h.GetUserByID)
		h.mux.HandleFunc("/users/add", h.AddUser)
		h.mux.HandleFunc("/users/delete/{id}", h.DeleteUser)
		h.mux.HandleFunc("/users/update/{id}", h.UpdateUser)
	}
}

// BookHandlers

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.ListBooks()
	if err != nil {
		log.Printf("GetBooks - h.service.ListBooks error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		log.Printf("GetBooks - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetBooks - data: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")

	if idStr == "" {
		log.Printf("GetBookByID - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBookByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.service.FindBook(id)
	if err != nil {
		log.Printf("GetBookByID - h.service.FindBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		log.Printf("GetBookByID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/books/author/")

	if idStr == "" {
		log.Printf("GetBooksByAuthor - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBooksByAuthor - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books, err := h.service.GetBooksByAuthor(id)
	if err != nil {
		log.Printf("GetBooksByAuthor - h.service.GetBooksByAuthor error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		log.Printf("GetBooksByAuthor - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddBook - incoming request: %v", string(data))

	var book model.Book

	err = json.Unmarshal(data, &book)
	if err != nil {
		log.Printf("AddBook - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddBook - data after unmarshalling: %v", book)

	createBook, err := h.service.CreateBook(&book)
	if err != nil {
		log.Printf("AddBook - h.service.CreateBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddBook - created book: %v", createBook)

	data, err = json.MarshalIndent(createBook, "", "    ")
	if err != nil {
		log.Printf("AddBook - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddBook - response to client: %v", string(data))

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("UpdateBook - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateBook - incoming request: %v", string(data))

	var book model.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		log.Printf("UpdateBook - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateBook - data after unmarshalling: %v", book)

	updatedBook, err := h.service.EditBook(&book)
	if err != nil {
		log.Printf("UpdateBook - h.service.EditBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateBook - updated book: %v", updatedBook)

	data, err = json.MarshalIndent(updatedBook, "", "    ")
	if err != nil {
		log.Printf("UpdateBook - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("UpdateBook - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/books/delete/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetBookByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.RemoveBook(id)
	if err != nil {
		log.Printf("GetBookByID - h.service.RemoveBook error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted"))
}

// UserHandlers

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

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("AddUser - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddUser - incoming request: %v", string(data))

	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Printf("AddUser - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddUser - data after unmarshalling: %v", user)

	createUser, err := h.service.CreateUser(&user)
	if err != nil {
		log.Printf("AddUser - h.service.CreateUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddUser - created user: %v", createUser)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(createUser, "", "    ")
	if err != nil {
		log.Printf("AddUser - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("AddUser - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
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

	_, err = h.service.RemoveUser(id)
	if err != nil {
		log.Printf("DeleteUser - h.service.RemoveUser error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
