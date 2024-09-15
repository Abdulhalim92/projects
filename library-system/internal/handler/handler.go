package handler

import (
	"net/http"
	"projects/internal/service"
)

type MyHandler struct {
	mux     *http.ServeMux
	service *service.Service
}

func NewMyHandler(mux *http.ServeMux, s *service.Service) *MyHandler {
	return &MyHandler{
		mux:     mux,
		service: s,
	}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *MyHandler) InitRoutes() {

	//Book Routes
	{
		h.mux.HandleFunc("/books", h.GetBooks)
		//h.mux.HandleFunc("/books/:id", h.GetBookByID)
		h.mux.HandleFunc("/books/{book_id}", h.GetBookByID)
		h.mux.HandleFunc("/books/authors/{author_id}", h.GetBooksByAuthor)
		h.mux.HandleFunc("/books/add", h.AddBook)
		h.mux.HandleFunc("/books/edit", h.EditBook)
		h.mux.HandleFunc("/books/delete/{id}", h.DeleteBook)
	}

	//User Routes
	{
		h.mux.HandleFunc("/users", h.GetUsers)
		h.mux.HandleFunc("/users/{user_id}", h.GetUserByID)
		h.mux.HandleFunc("/users/update", h.UpdateUser)
		h.mux.HandleFunc("/users/add", h.AddUser)
		h.mux.HandleFunc("/users/{user_id}/delete", h.DeleteUser)
	}

	//Authors Routes
	{
		h.mux.HandleFunc("/authors", h.GetAuthors)
		h.mux.HandleFunc("/authors/{id}", h.GetAuthorByID)
		h.mux.HandleFunc("/authors/add", h.CreateAuthor)
		h.mux.HandleFunc("/authors/update", h.UpdateAuthor)
	}

	//Borrows Routes
	{
		h.mux.HandleFunc("/borrows", h.GetBorrows)
		h.mux.HandleFunc("/borrows/filtered", h.GetFilteredBorrows)
		h.mux.HandleFunc("/borrows/add", h.AddBorrow)
		h.mux.HandleFunc("/borrows/return", h.ReturnBorrow)
	}

	//Review Routes
	{
		h.mux.HandleFunc("/reviews", h.GetReviews)
		h.mux.HandleFunc("/reviews/filtered", h.GetFilteredReviews)
		h.mux.HandleFunc("/reviews/add", h.AddReview)
	}
}
