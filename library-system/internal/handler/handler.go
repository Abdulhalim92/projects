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
	c.JSON(http.StatusOK, gin.H{"status": "The service is up and running"})
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
		bookGroup.GET("/:id", h.GetBookByID)
		bookGroup.POST("/add", h.AddBook)
		bookGroup.DELETE("/delete/:id", h.DeleteBook)
		bookGroup.PUT("/update", h.UpdateBook)
		bookGroup.GET("/author/:id", h.GetBooksByAuthor)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.GET("/:id", h.GetUserByID)
		userGroup.DELETE("/delete/:id", h.DeleteUser)
		userGroup.PUT("/update", h.UpdateUser)
	}

	authorGroup := v1.Group("/authors")
	{
		authorGroup.GET("", h.GetAuthors)
		authorGroup.GET("/:id", h.GetAuthorByID)
		authorGroup.POST("/add", h.AddAuthor)
		authorGroup.DELETE("/delete/:id", h.DeleteAuthor)
		authorGroup.PUT("/update", h.UpdateAuthor)
	}

	reviewGroup := v1.Group("/reviews")
	{
		reviewGroup.GET("", h.GetReviews)
		reviewGroup.GET("/:id", h.GetReviewByID)
		reviewGroup.GET("/user/:id", h.GetReviewsByUser)
		reviewGroup.GET("/book/:id", h.GetReviewsByBook)
		reviewGroup.GET("/filter", h.GetReviewsByFilter)
		reviewGroup.POST("/add", h.AddReview)
		reviewGroup.DELETE("/delete/:id", h.DeleteReview)
		reviewGroup.PUT("/update", h.UpdateReview)
	}

	profileGroup := v1.Group("/profiles")
	{
		profileGroup.GET("", h.ListProfiles)
		profileGroup.GET("/:id", h.GetProfileByID)
		profileGroup.POST("/add", h.CreateProfile)
		profileGroup.DELETE("/delete/:id", h.DeleteProfile)
		profileGroup.PUT("/update", h.UpdateProfile)
	}

	borrowGroup := v1.Group("/borrows")
	{
		borrowGroup.GET("", h.GetBorrows)
		borrowGroup.GET("/:id", h.GetBorrowByID)
		borrowGroup.GET("/user/:id", h.GetBorrowByUser)
		borrowGroup.GET("/book/:id", h.GetBorrowByBook)
		borrowGroup.POST("/add", h.CreateBorrow)
		borrowGroup.PUT("/return/:id", h.ReturnBook)
	}
}
