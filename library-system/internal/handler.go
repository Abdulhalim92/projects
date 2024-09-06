package internal

import (
	"net/http"
	"projects/internal/book"
	"projects/internal/user"
)

type MyHandler struct {
	mux *http.ServeMux
	book.BookHandler
	user.UserHandler
}

func NewMyHandler(mux *http.ServeMux, bookHandler book.BookHandler, userHandler user.UserHandler) *MyHandler {
	return &MyHandler{
		mux:         mux,
		BookHandler: bookHandler,
		UserHandler: userHandler,
	}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
