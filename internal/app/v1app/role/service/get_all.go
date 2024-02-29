package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.Role, error) {
	return s.roleRepository.GetAll(ctx, s.pagination.SetOffset(page))
}
