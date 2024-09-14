package internal

//
//import (
//	"net/http"
//	"projects/internal/handler"
//)
//
//type MyHandler struct {
//	mux *http.ServeMux
//	handler.BookHandler
//	handler.UserHandler
//}
//
//func NewMyHandler(mux *http.ServeMux, bookHandler handler.BookHandler, userHandler handler.UserHandler) *MyHandler {
//	return &MyHandler{
//		mux:         mux,
//		BookHandler: bookHandler,
//		UserHandler: userHandler,
//	}
//}
//
//func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	h.mux.ServeHTTP(w, r)
//}
