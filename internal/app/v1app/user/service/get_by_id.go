package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
)

const admin = true

// GetByID ...
func (s *service) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// todo add role
	// user.Role == "admin"
	if user.ID != id || !admin {
		return nil, errAccessDenied
	}

	return user, nil
}
