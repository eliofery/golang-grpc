package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
)

// SignIn ...
func (s *service) SignIn(ctx context.Context, user *dto.User) error {
	findUser, err := s.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if err = s.compareHashAndPassword(findUser.Password, user.Password); err != nil {
		return err
	}

	token, err := s.tokenManager.Generate(findUser.ID)
	if err != nil {
		return err
	}

	if err = s.tokenManager.SendAuthHeader(ctx, token); err != nil {
		return err
	}

	return nil
}
