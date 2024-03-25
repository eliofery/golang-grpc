package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, reqID int64) (*model.User, error) {
	user, err := s.accessManager.User(ctx)
	if err != nil {
		return nil, core.ErrAccessDenied
	}

	if reqID != user.ID && !s.accessManager.IsAccess(ctx, readPermission) {
		return nil, core.ErrAccessDenied
	}

	findUser, err := s.userRepository.GetByID(ctx, reqID)
	if err != nil {
		return nil, err
	}

	return findUser, nil
}
