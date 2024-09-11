package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"projects/internal/model"
	"projects/internal/service"
	"strconv"
)

type BookHandler struct {
	mux     *http.ServeMux
	service *service.BooksService
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func NewBookHandler(mux *http.ServeMux, s *service.BooksService) *BookHandler {
	return &BookHandler{
		mux:     mux,
		service: s,
	}
}

func (h *BookHandler) InitRoutes() {
	h.mux.HandleFunc("/books", h.GetBooks)
	//h.mux.HandleFunc("/books/:id", h.GetBookByID)
	h.mux.HandleFunc("/books/{book_id}", h.GetBookByID)
	h.mux.HandleFunc("/books/authors/{author_id}", h.GetBooksByAuthor)
	h.mux.HandleFunc("/books/add", h.AddBook)
	h.mux.HandleFunc("/books/edit", h.EditBook)
	h.mux.HandleFunc("/books/delete/{id}", h.DeleteBook)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
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

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
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

func (h *BookHandler) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
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

func (h *BookHandler) EditBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book model.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.service.EditBook(&book)

}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book model.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.service.RemoveBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
