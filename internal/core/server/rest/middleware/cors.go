package middleware

import (
	"net/http"
)

// cors ...
func cors(config *Config) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", config.Origin)
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Origin, Accept, Authorization, Content-Type, X-CSRF-Token")
			header.Set("Access-Control-Expose-Headers", "Content-Length, Grpc-Metadata-Authorization")
			header.Set("Access-Control-Max-Age", config.MaxAge)

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
