package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/core"
)

// Create ...
func (s *service) Create(ctx context.Context, user *dto.User) (int64, error) {
	if !s.accessManager.IsAccess(ctx, createPermission) {
		return 0, core.ErrAccessDenied
	}

	var err error
	user.Role.ID, err = s.settingRepository.GetDefaultRoleByID(ctx)
	if err != nil {
		return 0, err
	}

	user.Password, err = s.passwordManager.GenerateFromPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.ID, err = s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
