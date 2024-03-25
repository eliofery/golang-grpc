package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.User, error) {
	if !s.accessManager.IsAccess(ctx, readPermission) {
		return nil, core.ErrAccessDenied
	}

	users, err := s.userRepository.GetAll(ctx, s.pagination.SetOffset(page))
	if err != nil {
		return nil, err
	}

	return users, nil
}
