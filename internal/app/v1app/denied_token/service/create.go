package service

import "context"

// Create ...
func (s *service) Create(ctx context.Context, token string) error {
	return s.tokenRepository.Create(ctx, token)
}
