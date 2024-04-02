package middleware

import (
	"net/http"

	"github.com/eliofery/golang-grpc/pkg/eslog"
)

// Cors ...
type Cors interface {
	Cors() func(handler http.Handler) http.Handler
}

type cors struct {
	logger eslog.Logger
	config *Config
}

// NewCors ...
func NewCors(logger eslog.Logger, config *Config) Cors {
	return &cors{
		logger: logger,
		config: config,
	}
}

// Cors ...
func (c *cors) Cors() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", c.config.Origin)
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Origin, Accept, Authorization, Content-Type, X-CSRF-Token")
			header.Set("Access-Control-Expose-Headers", "Content-Length, Grpc-Metadata-Authorization")
			header.Set("Access-Control-Max-Age", c.config.MaxAge)

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
