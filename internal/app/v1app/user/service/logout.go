package service

import (
	"context"
)

// Logout ...
func (s *service) Logout(ctx context.Context, token string) error {
	return s.deniedTokenRepository.Create(ctx, token)
}
