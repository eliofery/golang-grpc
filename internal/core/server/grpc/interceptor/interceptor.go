package interceptor

import (
	deniedTokenRepository "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	rolePermissionRepository "github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// New ...
func New(
	logger *eslog.Logger,
	tokenManager *jwt.TokenManager,

	deniedTokenRepository deniedTokenRepository.Repository,
	rolePermissionRepository rolePermissionRepository.Repository,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		panicRecovery(logger),
		validate(logger),
		userDataFromToken(logger, tokenManager, deniedTokenRepository, rolePermissionRepository),
	}
}
