package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	role, err := s.roleRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return role, nil
}
