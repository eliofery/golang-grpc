package interceptor

import (
	"context"
	"log/slog"

	deniedTokenRepository "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	rolePermissionRepository "github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type key string

const userKey key = "user"

var (
	// ErrAlreadyAuthenticated ...
	ErrAlreadyAuthenticated = status.Error(codes.PermissionDenied, "you are already authenticated")

	// ErrNotAuthenticated ...
	ErrNotAuthenticated = status.Error(codes.PermissionDenied, "you are not authenticated")

	// ErrAccessDenied ...
	ErrAccessDenied = status.Error(codes.PermissionDenied, "access denied")
)

// UserData ...
type UserData struct {
	ID          int64
	Permissions []string
	Token       string
}

// userDataFromToken ...
func userDataFromToken(
	logger *eslog.Logger,
	tokenManager *jwt.TokenManager,

	deniedTokenRepository deniedTokenRepository.Repository,
	rolePermissionRepository rolePermissionRepository.Repository,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		op := "core.server.grpc.interceptor.userDataFromToken"

		token, err := tokenManager.GetAuthHeader(ctx)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		claims, err := tokenManager.Verify(token)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		if deniedToken, _ := deniedTokenRepository.GetByToken(ctx, token); deniedToken != nil {
			logger.Debug(op, slog.String("err", "denied token"))
			return handler(ctx, req)
		}

		userID, err := tokenManager.GetSubject(claims)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		roleID, err := tokenManager.GetRole(claims)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		permissions, err := rolePermissionRepository.GetPermissionsByRoleID(ctx, roleID)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		user := UserData{
			ID:          userID,
			Permissions: permissions,
			Token:       token,
		}

		return handler(withUser(ctx, &user), req)
	}
}

// withUser ...
func withUser(ctx context.Context, user *UserData) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User ...
func User(ctx context.Context) *UserData {
	user, ok := ctx.Value(userKey).(*UserData)
	if ok {
		return user
	}

	return nil
}
