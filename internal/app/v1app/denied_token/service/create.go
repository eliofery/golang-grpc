package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Create ...
func (s *service) Create(ctx context.Context, token string) error {
	if !interceptor.IsAccess(ctx, createPermission) {
		return interceptor.ErrAccessDenied
	}

	return s.deniedTokenRepository.Create(ctx, token)
}
