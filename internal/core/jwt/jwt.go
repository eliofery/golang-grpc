package jwt

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	roleKey       = "role"
	bearerPrefix  = "Bearer "
	authHeaderKey = "Authorization"
)

var (
	errCreateToken  = status.Error(codes.Internal, "failed create token")
	errGetToken     = status.Error(codes.Internal, "failed get token")
	errInvalidToken = status.Error(codes.InvalidArgument, "invalid token")
	errUserSubject  = status.Error(codes.Internal, "failed get user subject")
	errUserRole     = status.Error(codes.Internal, "failed get user role")
)

// TokenManager ...
type TokenManager struct {
	config *Config
	logger *eslog.Logger
}

// New ...
func New(config *Config, logger *eslog.Logger) *TokenManager {
	return &TokenManager{
		config: config,
		logger: logger,
	}
}

// Generate ...
func (t *TokenManager) Generate(userID, roleID int64) (string, error) {
	op := "core.jwt.Generate"

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   strconv.FormatInt(userID, 10),
		roleKey: strconv.FormatInt(roleID, 10),
		"exp":   time.Now().Add(time.Second * t.config.Expires).Unix(),
	})

	token, err := claims.SignedString([]byte(t.config.Secret))
	if err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return "", errCreateToken
	}

	return token, nil
}

// Verify ...
func (t *TokenManager) Verify(token string) (claims jwt.MapClaims, err error) {
	op := "core.jwt.Verify"

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.logger.Debug(op, slog.String("err", "token signing method must be: *jwt.SigningMethodHMAC"))
			return nil, errInvalidToken
		}

		return []byte(t.config.Secret), nil
	})

	if err != nil || !parsedToken.Valid {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.logger.Debug(op, slog.String("err", "token claims must be: jwt.MapClaims"))
		return nil, errInvalidToken
	}

	return claims, nil
}

// GetSubject ...
func (t *TokenManager) GetSubject(claims jwt.MapClaims) (int64, error) {
	op := "core.jwt.GetSubject"

	id, err := claims.GetSubject()
	if err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return 0, errUserSubject
	}

	return strconv.ParseInt(id, 10, 64)
}

// GetRole ...
func (t *TokenManager) GetRole(claims jwt.MapClaims) (int64, error) {
	op := "core.jwt.GetRole"

	id, ok := claims[roleKey].(string)
	if !ok {
		t.logger.Debug(op, slog.String("err", fmt.Sprintf("%s is invalid", roleKey)))
		return 0, errUserRole
	}

	return strconv.ParseInt(id, 10, 64)
}

// SendAuthHeader ...
func (t *TokenManager) SendAuthHeader(ctx context.Context, token string) error {
	op := "core.jwt.SendAuthHeader"

	if err := grpc.SendHeader(ctx, metadata.Pairs(authHeaderKey, bearerPrefix+token)); err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return errGetToken
	}

	return nil
}

// GetAuthHeader ...
func (t *TokenManager) GetAuthHeader(ctx context.Context) (string, error) {
	op := "core.jwt.GetAuthHeader"

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		t.logger.Debug(op, slog.String("err", "metadata not provided"))
		return "", errGetToken
	}

	authHeader := md.Get(authHeaderKey)
	if authHeader == nil {
		t.logger.Debug(op, slog.String("err", "authorization header not provided"))
		return "", errGetToken
	}

	if !strings.HasPrefix(authHeader[0], bearerPrefix) {
		t.logger.Debug(op, slog.String("err", "invalid authorization header format"))
		return "", errGetToken
	}

	return strings.TrimPrefix(authHeader[0], bearerPrefix), nil
}
