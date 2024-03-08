package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.User, error) {
	return s.userRepository.GetAll(ctx, s.pagination.SetOffset(page))
}
