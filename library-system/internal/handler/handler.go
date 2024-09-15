package handler

import (
	"net/http"
	"projects/internal/handler/middleware"
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
	publicMiddleware := []middleware.Middleware{
		middleware.CORS,
		middleware.Recovery,
	}

	protectedMiddleware := append(publicMiddleware, middleware.Authenticate)

	{
		h.mux.Handle("/health", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.HealthCheck, http.MethodGet), publicMiddleware...))
		h.mux.Handle("/users/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.SignUp, http.MethodPost), publicMiddleware...))
		h.mux.Handle("/users/login", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.SignIn, http.MethodPost), publicMiddleware...))
	}

	bookGroup := h.Group("/books")
	{
		bookGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBooks, http.MethodGet), protectedMiddleware...))
		bookGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBookByID, http.MethodGet), protectedMiddleware...))
		bookGroup.Handle("/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.AddBook, http.MethodGet), protectedMiddleware...))
		bookGroup.Handle("/delete/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.DeleteBook, http.MethodGet), protectedMiddleware...))
		bookGroup.Handle("/update", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.UpdateBook, http.MethodGet), protectedMiddleware...))
		bookGroup.Handle("/author/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBooksByAuthor, http.MethodGet), protectedMiddleware...))
	}

	userGroup := h.Group("/users")
	{
		userGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetUsers, http.MethodGet), protectedMiddleware...))
		userGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetUserByID, http.MethodGet), protectedMiddleware...))
		userGroup.Handle("/delete/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.DeleteUser, http.MethodGet), protectedMiddleware...))
		userGroup.Handle("/update", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.UpdateUser, http.MethodGet), protectedMiddleware...))
	}

	authorGroup := h.Group("/authors")
	{
		authorGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetAuthors, http.MethodGet), protectedMiddleware...))
		authorGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetAuthorByID, http.MethodGet), protectedMiddleware...))
		authorGroup.Handle("/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.AddAuthor, http.MethodGet), protectedMiddleware...))
		authorGroup.Handle("/delete/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.DeleteAuthor, http.MethodGet), protectedMiddleware...))
		authorGroup.Handle("/update", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.UpdateAuthor, http.MethodGet), protectedMiddleware...))
	}

	reviewGroup := h.Group("/reviews")
	{
		reviewGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetReviews, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetReviewByID, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/user/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetReviewsByUser, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/book/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetReviewsByBook, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/filter", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetReviewsByFilter, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.AddReview, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/delete/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.DeleteReview, http.MethodGet), protectedMiddleware...))
		reviewGroup.Handle("/update", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.UpdateReview, http.MethodGet), protectedMiddleware...))
	}

	profileGroup := h.Group("/profiles")
	{
		profileGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.ListProfiles, http.MethodGet), protectedMiddleware...))
		profileGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetProfileByID, http.MethodGet), protectedMiddleware...))
		profileGroup.Handle("/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.CreateProfile, http.MethodGet), protectedMiddleware...))
		profileGroup.Handle("/delete/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.DeleteProfile, http.MethodGet), protectedMiddleware...))
		profileGroup.Handle("/update}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.UpdateProfile, http.MethodGet), protectedMiddleware...))
	}

	borrowGroup := h.Group("/borrows")
	{
		borrowGroup.Handle("", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBorrows, http.MethodGet), protectedMiddleware...))
		borrowGroup.Handle("/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBorrowByID, http.MethodGet), protectedMiddleware...))
		borrowGroup.Handle("/user/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBorrowByUser, http.MethodGet), protectedMiddleware...))
		borrowGroup.Handle("/book/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.GetBorrowByBook, http.MethodGet), protectedMiddleware...))
		borrowGroup.Handle("/add", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.CreateBorrow, http.MethodPost), protectedMiddleware...))
		borrowGroup.Handle("/return/{id}", middleware.ChainMiddleware(middleware.MethodCheckHandler(h.ReturnBook, http.MethodGet), protectedMiddleware...))
	}
}
