package middleware

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// New ...
func New(
	mux *runtime.ServeMux,
	cors Cors,
) http.Handler {
	return chainMiddlewares(
		mux,
		cors.Cors(),
	)
}

// chainMiddlewares ...
func chainMiddlewares(mux *runtime.ServeMux, handlers ...func(http.Handler) http.Handler) http.Handler {
	if len(handlers) == 0 {
		return mux
	}

	var chainMiddleware http.Handler = mux
	for _, handler := range handlers {
		chainMiddleware = handler(chainMiddleware)
	}

	return chainMiddleware
}
