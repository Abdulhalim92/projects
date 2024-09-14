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

func (h *MyHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.ListBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	//path := strings.Split(r.URL.RawQuery, "/")
	//id := path[len(path)-1]
	id, err := strconv.Atoi(r.PathValue("book_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book, err := h.service.FindBook(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(book, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := strconv.Atoi(r.PathValue("author_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books, err := h.service.FindBooksByAuthor(authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(books, "   ", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (h *MyHandler) EditBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book *model.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book, err = h.service.EditBook(book)
	if err != nil {
		log.Printf("Failed to Edit Book - service.EditBook error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(book, "   ", "")
	if err != nil {
		log.Printf("Failed to Edit Book - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Edit Book - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *MyHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book *model.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book, err = h.service.CreateBook(book)
	if err != nil {
		log.Printf("Failed to Add Book - service.CreateBook error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err = json.MarshalIndent(book, "   ", "")
	if err != nil {
		log.Printf("Failed to Add Book - json.MarshalIndent error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to Add Book - Write error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MyHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err = h.service.RemoveBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Book with id %d was successfully deleted", id)
}
