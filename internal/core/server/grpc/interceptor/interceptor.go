package interceptor

import (
	"github.com/eliofery/golang-fullstack/internal/core/jwt"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"google.golang.org/grpc"
)

// New ...
func New(logger *eslog.Logger, tokenManager *jwt.TokenManager) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		validate(logger),
		userIDFromToken(logger, tokenManager),
	}
}
