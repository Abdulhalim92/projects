package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/internal/handler/middleware"
	"projects/internal/service"
)

type Handler struct {
	router  *gin.Engine
	service *service.Service
}

func NewHandler(router *gin.Engine, s *service.Service) *Handler {
	return &Handler{
		router:  router,
		service: s,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "The service is up and running")
}

func (h *Handler) InitRoutes() {
	h.router.Use(
		middleware.CORS(),
		middleware.Recovery(),
	)

	{
		h.router.GET("/health", h.HealthCheck)
		h.router.POST("/users/add", h.SignUp)
		h.router.POST("/users/login", h.SignIn)
	}

	v1 := h.router.Group("/v1")
	v1.Use(
		middleware.Authenticate(),
	)

	bookGroup := v1.Group("/books")
	{
		bookGroup.GET("", h.GetBooks)
		bookGroup.GET("/{id}", h.GetBookByID)
		bookGroup.POST("/add", h.AddBook)
		bookGroup.DELETE("/delete/{id}", h.DeleteBook)
		bookGroup.PUT("/update", h.UpdateBook)
		bookGroup.GET("/author/{id}", h.GetBooksByAuthor)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.GET("/{id}", h.GetUserByID)
		userGroup.GET("/delete/{id}", h.DeleteUser)
		userGroup.GET("/update", h.UpdateUser)
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
