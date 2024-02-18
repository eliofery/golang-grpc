package middleware

import (
	"net/http"

	"github.com/eliofery/golang-fullstack/pkg/esmux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// ChainMiddlewares ...
func ChainMiddlewares(mux *runtime.ServeMux) http.Handler {
	return esmux.ChainMiddlewares(mux)(
		CORS,
	)
}
