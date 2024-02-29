package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
)

// Update ...
func (s *service) Update(ctx context.Context, role *dto.Role) (*model.Role, error) {
	roleUpdate, err := s.roleRepository.Update(ctx, role)
	if err != nil {
		return nil, err
	}

	return roleUpdate, nil
}
