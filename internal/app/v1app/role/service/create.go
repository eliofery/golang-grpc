package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
)

func (s *service) Create(ctx context.Context, role *dto.Role) (int64, error) {
	return s.roleRepository.Create(ctx, role)
}
