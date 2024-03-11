package interceptor

import (
	"context"
	"log/slog"

	deniedTokenV1Repository "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	userV1Repository "github.com/eliofery/golang-grpc/internal/app/v1app/user/repository"
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
)

// UserData ...
type UserData struct {
	User  *model.User
	Token string
}

// userDataFromAuthHeader ...
func userDataFromAuthHeader(
	logger *eslog.Logger,
	tokenManager *jwt.TokenManager,

	deniedTokenV1Repository deniedTokenV1Repository.Repository,
	userV1Repository userV1Repository.Repository,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		op := "core.server.grpc.interceptor.userDataFromAuthHeader"

		token, err := tokenManager.GetAuthHeader(ctx)
		if err != nil {
			return handler(ctx, req)
		}

		claims, err := tokenManager.Verify(token)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		if deniedToken, _ := deniedTokenV1Repository.GetByToken(ctx, token); deniedToken != nil {
			logger.Debug(op, slog.String("err", "denied token"))
			return handler(ctx, req)
		}

		userID, err := tokenManager.GetSubject(claims)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		user, err := userV1Repository.GetByID(ctx, userID)
		if err != nil {
			logger.Debug(op, slog.String("err", err.Error()))
			return handler(ctx, req)
		}

		userData := UserData{
			User:  user,
			Token: token,
		}

		return handler(withUser(ctx, &userData), req)
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

// UserID ...
func UserID(ctx context.Context, reqID ...int64) int64 {
	if len(reqID) > 0 && reqID[0] != 0 {
		return reqID[0]
	}

	return User(ctx).User.ID
}

// UserToken ...
func UserToken(ctx context.Context) string {
	return User(ctx).Token
}

// UserRoleID ...
func UserRoleID(ctx context.Context) int64 {
	return User(ctx).User.Role.ID
}

// IsAuthenticated ...
func IsAuthenticated(ctx context.Context) bool {
	return User(ctx) != nil
}
