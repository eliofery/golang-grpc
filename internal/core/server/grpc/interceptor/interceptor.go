package interceptor

import (
	deniedtokenv1 "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	userV1Repository "github.com/eliofery/golang-grpc/internal/app/v1app/user/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// New ...
func New(
	logger *eslog.Logger,
	tokenManager *jwt.TokenManager,

	deniedtokenv1 deniedtokenv1.Repository,
	userV1Repository userV1Repository.Repository,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		panicRecovery(logger),
		validate(logger),
		userDataFromAuthHeader(logger, tokenManager, deniedtokenv1, userV1Repository),
	}
}
