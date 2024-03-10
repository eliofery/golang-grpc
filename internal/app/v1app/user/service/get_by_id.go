package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
