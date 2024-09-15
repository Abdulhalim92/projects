package handler

import (
	"net/http"
	"projects/internal/handler/middleware"
)

// RouteGroup структура для хранения группы маршрутов.
type RouteGroup struct {
	prefix     string
	mux        *http.ServeMux
	middleware []func(http.Handler) http.Handler
}

// Group создает новую группу маршрутов с общим префиксом и middleware.
func (h *Handler) Group(prefix string, middleware ...func(http.Handler) http.Handler) *RouteGroup {
	return &RouteGroup{
		prefix:     prefix,
		mux:        h.mux,
		middleware: middleware,
	}
}

// Handle добавляет маршрут в группу маршрутов с учетом middleware.
func (rg *RouteGroup) Handle(path string, handler http.Handler) {
	fullPath := rg.prefix + path // Комбинируем префикс группы с путем

	// Примеряем все middleware в маршрутизатор
	finalHandler := handler
	for _, mw := range rg.middleware {
		finalHandler = mw(finalHandler)
	}

	// Добавляем обработчик с учетом middleware в маршрутизатор
	rg.mux.Handle(fullPath, finalHandler)
}

func (rg *RouteGroup) HandleV2(path string, handler http.Handler, allowedMethods ...string) {
	fullPath := rg.prefix + path // Комбинируем префикс группы с путем

	// Применяем middleware для проверки метода запроса
	finalHandler := middleware.MethodCheckMiddleware(handler, allowedMethods...)

	// Применяем все остальные middleware к обработчику
	for _, mw := range rg.middleware {
		finalHandler = mw(finalHandler)
	}

	// Добавляем обработчик с учетом всех middleware в маршрутизатор
	rg.mux.Handle(fullPath, finalHandler)
}
