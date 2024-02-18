package esmux

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// UseMiddlewares ...
type UseMiddlewares func(...func(http.Handler) http.Handler) http.Handler

// ChainMiddlewares ...
func ChainMiddlewares(mux *runtime.ServeMux) UseMiddlewares {
	return func(handlers ...func(http.Handler) http.Handler) http.Handler {
		if len(handlers) == 0 {
			return mux
		}

		var chainMiddleware http.Handler
		for key, handler := range handlers {
			if key == 0 {
				chainMiddleware = handler(mux)
				continue
			}

			chainMiddleware = handler(chainMiddleware)
		}

		return chainMiddleware
	}
}
