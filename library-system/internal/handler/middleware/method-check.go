package middleware

import "net/http"

// MethodCheckHandler обертка для проверки метода.
func MethodCheckHandler(h http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		methodAllowed := false
		for _, method := range allowedMethods {
			if r.Method == method {
				methodAllowed = true
				break
			}
		}

		if !methodAllowed {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Если метод разрешен, вызываем обработчик
		h.ServeHTTP(w, r)
	}
}

// Пример middleware для проверки метода запроса.
func MethodCheckMiddleware(handler http.Handler, allowedMethods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		methodAllowed := false
		for _, method := range allowedMethods {
			if r.Method == method {
				methodAllowed = true
				break
			}
		}

		if !methodAllowed {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Если метод разрешен, вызываем обработчик
		handler.ServeHTTP(w, r)
	})
}
