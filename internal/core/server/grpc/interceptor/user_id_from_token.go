package interceptor

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/eliofery/golang-fullstack/internal/core/jwt"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type key string

const (
	// UserIDKey ...
	UserIDKey key = "userID"

	// BearerToken ...
	BearerToken = "Bearer "
)

var (
	// ErrAlreadyAuthenticated ...
	ErrAlreadyAuthenticated = status.Error(codes.PermissionDenied, "you are already authenticated")

	// ErrNotAuthenticated ...
	ErrNotAuthenticated = status.Error(codes.PermissionDenied, "you are not authenticated")
)

// userIDFromToken ...
func userIDFromToken(logger *eslog.Logger, tokenManager *jwt.TokenManager) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		op := "core.server.grpc.interceptor.userIDFromToken"

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Debug(op, slog.String("err", "metadata not provided"))
			return handler(ctx, req)
		}

		authHeader := md.Get("Authorization")
		if authHeader == nil {
			logger.Debug(op, slog.String("err", "authorization header not provided"))
			return handler(ctx, req)
		}

		if !strings.HasPrefix(authHeader[0], BearerToken) {
			logger.Debug(op, slog.String("err", "invalid authorization header format"))
			return handler(ctx, req)
		}

		token := strings.TrimPrefix(authHeader[0], BearerToken)

		claims, err := tokenManager.Verify(token)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		userID, err := tokenManager.GetSubject(claims)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		return handler(withUserID(ctx, userID), req)
	}
}

// withUserID ...
func withUserID(ctx context.Context, userID *int64) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// userID ...
func userID(ctx context.Context) *int64 {
	id, ok := ctx.Value(UserIDKey).(*int64)
	if ok {
		return id
	}

	return nil
}

// IsAuthenticated ...
func IsAuthenticated(ctx context.Context) bool {
	return userID(ctx) != nil
}
