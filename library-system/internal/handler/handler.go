package handler

import (
	"net/http"
	"projects/internal/service"
)

type Handler struct {
	mux     *http.ServeMux
	service *service.Service
}

func NewHandler(mux *http.ServeMux, s *service.Service) *Handler {
	return &Handler{
		mux:     mux,
		service: s,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The service is up and running"))
}

func (h *Handler) InitRoutes() {
	h.mux.HandleFunc("/health", h.HealthCheck)
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
		h.mux.HandleFunc("/users/login", h.SignIn)
		h.mux.HandleFunc("/users/delete/{id}", h.DeleteUser)
		h.mux.HandleFunc("/users/update/{id}", h.UpdateUser)
	}
	{
		h.mux.HandleFunc("/authors", h.GetAuthors)
		h.mux.HandleFunc("/authors/{id}", h.GetAuthorByID)
		h.mux.HandleFunc("/authors/add", h.AddAuthor)
		h.mux.HandleFunc("/authors/delete/{id}", h.DeleteAuthor)
		h.mux.HandleFunc("/authors/update/{id}", h.UpdateAuthor)
	}
	{
		h.mux.HandleFunc("/reviews", h.GetReviews)
		h.mux.HandleFunc("/reviews/{id}", h.GetReviewByID)
		h.mux.HandleFunc("/reviews/user/{id}", h.GetReviewsByUser)
		h.mux.HandleFunc("/reviews/book/{id}", h.GetReviewsByBook)
		h.mux.HandleFunc("/reviews/filter", h.GetReviewsByFilter)
		h.mux.HandleFunc("/reviews/add", h.AddReview)
		h.mux.HandleFunc("/reviews/delete/{id}", h.DeleteReview)
		h.mux.HandleFunc("/reviews/update/{id}", h.UpdateReview)
	}
	{
		h.mux.HandleFunc("/profiles", h.ListProfiles)
		h.mux.HandleFunc("/profiles/{id}", h.GetProfileByID)
		h.mux.HandleFunc("/profiles/add", h.CreateProfile)
		h.mux.HandleFunc("/profiles/delete/{id}", h.DeleteProfile)
		h.mux.HandleFunc("/profiles/update/{id}", h.UpdateProfile)
	}
	{
		h.mux.HandleFunc("/borrows", h.GetBorrows)
		h.mux.HandleFunc("/borrows/{id}", h.GetBorrowByID)
		h.mux.HandleFunc("/borrows/user/{id}", h.GetBorrowByUser)
		h.mux.HandleFunc("/borrows/book/{id}", h.GetBorrowByBook)
		h.mux.HandleFunc("/borrows/add", h.CreateBorrow)
		h.mux.HandleFunc("/borrows/return/{id}", h.ReturnBook)
	}
}
