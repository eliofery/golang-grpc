package interceptor

import (
	"context"
	"log/slog"

	authRepository "github.com/eliofery/golang-grpc/internal/app/v1/repository/auth"
	"github.com/eliofery/golang-grpc/internal/core/authorize/access"
	"github.com/eliofery/golang-grpc/internal/core/authorize/token"
	"github.com/eliofery/golang-grpc/internal/core/metadata"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// Auth ...
type Auth interface {
	Authorize() grpc.UnaryServerInterceptor
}

type auth struct {
	logger          eslog.Logger
	tokenManager    token.Manager
	accessManager   access.Manager
	metadataManager metadata.Manager
	authRepository  authRepository.Repository
}

// NewAuth ...
func NewAuth(
	logger eslog.Logger,
	tokenManager token.Manager,
	accessManager access.Manager,
	metadataManager metadata.Manager,
	authRepository authRepository.Repository,
) Auth {
	return &auth{
		logger:          logger,
		tokenManager:    tokenManager,
		accessManager:   accessManager,
		metadataManager: metadataManager,
		authRepository:  authRepository,
	}
}

// Authorize ...
func (a *auth) Authorize() grpc.UnaryServerInterceptor {
	op := "core.server.grpc.interceptor.Authorize"

	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		jwtToken, err := a.metadataManager.GetAuthHeader(ctx)
		if err != nil {
			a.logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		userID, err := a.tokenManager.AccessTokenParse(jwtToken)
		if err != nil {
			a.logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		cacheUser, err := a.authRepository.GetUserCache(ctx, userID)
		if err != nil {
			a.logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		ctx = a.accessManager.WithUser(ctx, &access.User{
			ID:          cacheUser.ID,
			Permissions: cacheUser.Permissions,
		})

		return handler(ctx, req)
	}
}
