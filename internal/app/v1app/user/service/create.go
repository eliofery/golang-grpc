package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Create ...
func (s *service) Create(ctx context.Context, user *dto.User) (int64, error) {
	if !interceptor.IsAccess(ctx, createPermission) {
		return 0, interceptor.ErrAccessDenied
	}

	var err error
	user.RoleID, err = s.settingRepository.GetDefaultRoleID(ctx)
	if err != nil {
		return 0, err
	}

	user.Password, err = s.generateFromPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.ID, err = s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
