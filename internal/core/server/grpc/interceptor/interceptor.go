package interceptor

import (
	tokenv1 "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// New ...
func New(
	logger *eslog.Logger,
	tokenManager *jwt.TokenManager,
	tokenRepository tokenv1.Repository,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		panicRecovery(logger),
		validate(logger),
		userDataFromAuthHeader(logger, tokenManager, tokenRepository),
	}
}
