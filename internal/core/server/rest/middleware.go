package rest

import (
	"net/http"

	"github.com/eliofery/golang-fullstack/pkg/esmux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// NewMiddleware ...
func NewMiddleware(mux *runtime.ServeMux, config *Config) http.Handler {
	return esmux.ChainMiddlewares(
		mux,
		cors(config),
	)
}

// cors ...
func cors(config *Config) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", config.Origin)
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Origin, Accept, Authorization, Content-Type, X-CSRF-Token")
			header.Set("Access-Control-Expose-Headers", "Link, Content-Length, Access-Control-Allow-Origin")
			header.Set("Access-Control-Max-Age", config.MaxAge)

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
