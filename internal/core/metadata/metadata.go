package metadata

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	authDefaultPrefix = "Bearer"
	authHeaderKey     = "Authorization"
)

// Manager ...
type Manager interface {
	SendAuthHeader(ctx context.Context, token string, prefix ...string) error
	GetAuthHeader(ctx context.Context, prefix ...string) (string, error)
	GetIPAddress(ctx context.Context) (string, error)
	GetUserAgent(ctx context.Context) (string, error)
}

// metadata ...
type metadata struct {
	logger eslog.Logger
}

// New ...
func New(logger eslog.Logger) Manager {
	return &metadata{
		logger: logger,
	}
}

// SendAuthHeader ...
func (m *metadata) SendAuthHeader(ctx context.Context, token string, prefix ...string) error {
	op := "core.metadata.SendAuthHeader"

	if len(prefix) == 0 {
		prefix = []string{authDefaultPrefix}
	}

	token = fmt.Sprintf("%s %s", prefix[0], token)
	if err := grpc.SendHeader(ctx, grpcMetadata.Pairs(authHeaderKey, token)); err != nil {
		m.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrInternal
	}

	return nil
}

// GetAuthHeader ...
func (m *metadata) GetAuthHeader(ctx context.Context, prefix ...string) (string, error) {
	op := "core.metadata.GetAuthHeader"

	if len(prefix) == 0 {
		prefix = []string{authDefaultPrefix}
	}

	md, ok := grpcMetadata.FromIncomingContext(ctx)
	if !ok {
		m.logger.Debug(op, slog.String("err", "metadata not provided"))
		return "", core.ErrInternal
	}

	authHeader := md.Get(authHeaderKey)
	if authHeader == nil {
		m.logger.Debug(op, slog.String("err", "authorization header not provided"))
		return "", core.ErrInternal
	}

	if !strings.HasPrefix(authHeader[0], prefix[0]) {
		m.logger.Debug(op, slog.String("err", "invalid authorization header format"))
		return "", core.ErrInternal
	}

	return strings.TrimPrefix(strings.TrimPrefix(authHeader[0], prefix[0]), " "), nil
}

// GetIPAddress ...
func (m *metadata) GetIPAddress(ctx context.Context) (string, error) {
	op := "core.metadata.GetIPAddress"

	p, ok := peer.FromContext(ctx)
	if !ok {
		m.logger.Debug(op, slog.String("err", "cannot get peer from context"))
		return "", core.ErrInternal
	}

	return p.Addr.String(), nil
}

// GetUserAgent ...
func (m *metadata) GetUserAgent(ctx context.Context) (string, error) {
	op := "core.metadata.GetUserAgent"

	md, ok := grpcMetadata.FromIncomingContext(ctx)
	if !ok {
		m.logger.Debug(op, slog.String("err", "cannot get metadata from context"))
		return "", core.ErrInternal
	}

	return md.Get("user-agent")[0], nil
}
