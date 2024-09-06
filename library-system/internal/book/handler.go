package book

import (
	"encoding/json"
	"fmt"
	"io"
	"library-system/internal/model"
	"net/http"
	"strconv"
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
		fmt.Printf("error when marshlling books: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

	defer r.Body.Close()
	io.ReadAll(r.Body)

}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	bookID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	book, err := h.bookService.FindBook(bookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonBook, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Printf("error when marshlling book: %e", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
} // TODO: implement

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	title := r.FormValue("title")
	authorID, err := strconv.Atoi(r.FormValue("author_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var book = &model.Book{
		Title:    title,
		AuthorId: authorID,
	}

	book, err = h.bookService.br.AddBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBook, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)

} // TODO: implement

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	bookID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result := h.bookService.RemoveBook(bookID)
	if result == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Book with id:%d removed successfully", bookID)
	w.Write([]byte(response))
} // TODO: implement
