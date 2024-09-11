package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
	"strings"
)

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

	books, err := h.service.FindBooksByAuthor(id)
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

	//log.Printf("AddBook - incoming request: %v\n", string(data))

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
