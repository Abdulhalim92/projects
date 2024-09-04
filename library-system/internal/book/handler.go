package book

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BookHandler struct {
	mux         *http.ServeMux
	bookService *Service
}

func NewBookHandler(mux *http.ServeMux, s *Service) *BookHandler {
	return &BookHandler{
		mux:         mux,
		bookService: s,
	}
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// InitRoutes adds routes for books to the given ServeMux.
//
// It adds the following routes:
//
// - GET /books: GetBooks
// - GET /books/:id: GetBookByID
// - POST /books/add: AddBook
// - DELETE /books/delete/:id: DeleteBook
func (h *BookHandler) InitRoutes() {
	h.mux.HandleFunc("/books", h.GetBooks)
	h.mux.HandleFunc("/books/:id", h.GetBookByID)
	h.mux.HandleFunc("/books/add", h.AddBook)
	h.mux.HandleFunc("/books/delete/:id", h.DeleteBook)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.ListBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
} // TODO: implement

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
} // TODO: implement

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
} // TODO: implement
