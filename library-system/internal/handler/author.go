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

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAuthors()
	if err != nil {
		log.Printf("GetAuthors - h.service.ListAuthors error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetAuthors - authors: %v", authors)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(authors, "", "    ")
	if err != nil {
		log.Printf("GetAuthors - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetAuthors - data: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/authors/")

	if idStr == "" {
		log.Printf("GetAuthorByID - id is required")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetAuthorByID - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		log.Printf("GetAuthorByID - h.service.GetAuthorByID error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetAuthorByID - author: %v", author)

	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(author, "", "    ")
	if err != nil {
		log.Printf("GetAuthorByID - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("GetAuthorByID - data: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("CreateAuthor - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateAuthor - incoming request: %v", string(data))

	var author model.Author
	err = json.Unmarshal(data, &author)
	if err != nil {
		log.Printf("CreateAuthor - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateAuthor - data after unmarshalling: %v", author)

	createAuthor, err := h.service.CreateAuthor(&author)
	if err != nil {
		log.Printf("CreateAuthor - h.service.CreateAuthor error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateAuthor - created author: %v", createAuthor)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(createAuthor, "", "    ")
	if err != nil {
		log.Printf("CreateAuthor - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("CreateAuthor - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("EditAuthor - io.ReadAll error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditAuthor - incoming request: %v", string(data))

	var author model.Author
	err = json.Unmarshal(data, &author)
	if err != nil {
		log.Printf("EditAuthor - json.Unmarshal error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditAuthor - data after unmarshalling: %v", author)

	updateAuthor, err := h.service.EditAuthor(&author)
	if err != nil {
		log.Printf("EditAuthor - h.service.EditAuthor error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditAuthor - updated author: %v", updateAuthor)

	w.Header().Set("Content-Type", "application/json")

	data, err = json.MarshalIndent(updateAuthor, "", "    ")
	if err != nil {
		log.Printf("EditAuthor - json.MarshalIndent error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("EditAuthor - response to client: %v", string(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/authors/delete/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteAuthor - strconv.Atoi error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.service.DeleteAuthor(id)
	if err != nil {
		log.Printf("DeleteAuthor - h.service.DeleteAuthor error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
