package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// ChainMiddleware - функция для объединения нескольких middleware
func ChainMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
