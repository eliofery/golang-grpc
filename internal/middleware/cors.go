package middleware

import (
	"net/http"
)

// CORS ...
func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Origin, Accept, Authorization, Content-Type, X-CSRF-Token")
		header.Set("Access-Control-Expose-Headers", "Link, Content-Length, Access-Control-Allow-Origin")
		header.Set("Access-Control-Max-Age", "0")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
