package rest

import (
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
)

// Middleware ...
type Middleware struct {
	fx.Out

	Middlewares http.Handler
}

// NewMiddleware ...
func NewMiddleware(mux *runtime.ServeMux) Middleware {
	useMiddlewares := chainMiddlewares(
		mux,
		middleware.CORS,
	)

	return Middleware{
		Middlewares: useMiddlewares,
	}
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
