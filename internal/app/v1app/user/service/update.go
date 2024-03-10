package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
)

// Update ...
func (s *service) Update(ctx context.Context, user *dto.Update) (*model.User, error) {
	findUser, err := s.userRepository.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if user.OldPassword.Valid && user.NewPassword.Valid {
		if err = s.compareHashAndPassword(findUser.Password, user.OldPassword.String); err != nil {
			return nil, errWrongOldPassword
		}

		user.NewPassword.String, err = s.generateFromPassword(user.NewPassword.String)
		if err != nil {
			return nil, err
		}
	}

	updateUser, err := s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepository.GetByID(ctx, updateUser.Role.ID)
	if err != nil {
		return nil, err
	}
	updateUser.Role.Name = role.Name

	return updateUser, nil
}