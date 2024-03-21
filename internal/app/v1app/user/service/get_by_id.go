package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, reqID, userID int64) (*model.User, error) {
	if reqID != userID && !interceptor.IsAccess(ctx, readPermission) {
		return nil, interceptor.ErrAccessDenied
	}

	user, err := s.userRepository.GetByID(ctx, reqID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
