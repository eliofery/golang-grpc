package esmux

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// ChainMiddlewares ...
func ChainMiddlewares(mux *runtime.ServeMux, handlers ...func(http.Handler) http.Handler) http.Handler {
	if len(handlers) == 0 {
		return mux
	}

	var chainMiddleware http.Handler = mux
	for _, handler := range handlers {
		chainMiddleware = handler(chainMiddleware)
	}

	return chainMiddleware
}
