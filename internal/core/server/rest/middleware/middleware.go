package middleware

import (
	"net/http"

	"github.com/eliofery/golang-fullstack/pkg/esmux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// New ...
func New(mux *runtime.ServeMux, config *Config) http.Handler {
	return esmux.ChainMiddlewares(
		mux,
		cors(config),
	)
}
