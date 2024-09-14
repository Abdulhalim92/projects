package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/internal/model"
	"strconv"
)

func (h *MyHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.ListAuthors()
	if err != nil {
		log.Printf("Service List Authors error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(authors, "   ", "")
	if err != nil {
		log.Printf("Service List Authors error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Service List Authors error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (h *MyHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Failed to GetAuthorByID - strconv.Atoi error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	author, err := h.service.FindAuthorByID(id)
	if err != nil {
		log.Printf("Failed to GetAuthorByID - service.FindAuthorByID error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(author, "   ", "")
	if err != nil {
		log.Printf("Failed to GetAuthorByID - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to GetAuthorByID - Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (h *MyHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Create Author - io.ReadAll error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	var author *model.Author
	err = json.Unmarshal(data, &author)
	if err != nil {
		log.Printf("Failed to Create Author - json.Unmarshal error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	author, err = h.service.CreateAuthor(author)
	if err != nil {
		log.Printf("Failed to Create Author - service.CreateAuthor error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(author, "   ", "")
	if err != nil {
		log.Printf("Failed to Create Author - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Create Author - Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (h *MyHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to Update Author - io.ReadAll error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	var author *model.Author
	err = json.Unmarshal(data, &author)
	if err != nil {
		log.Printf("Failed to Update Author - json.Unmarshal error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	author, err = h.service.UpdateAuthor(author)
	if err != nil {
		log.Printf("Failed to Update Author - service.UpdateAuthor error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(author, "   ", "")
	if err != nil {
		log.Printf("Failed to Update Author - json.MarshalIndent error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Update Author - Write error %v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
