package service

import (
	"context"
)

// SignUp ...
func (s *service) SignUp(ctx context.Context, userID int64) error {
	token, err := s.tokenManager.Generate(userID)
	if err != nil {
		return err
	}

	if err = s.tokenManager.SendAuthHeader(ctx, token); err != nil {
		return err
	}

	return nil
}
