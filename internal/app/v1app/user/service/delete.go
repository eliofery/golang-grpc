package service

import (
	"context"
)

// Delete ...
func (s *service) Delete(ctx context.Context, id int64) error {
	if err := s.userRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
