package service

import (
	"context"
)

// Delete ...
func (s *service) Delete(ctx context.Context, id int64) error {
	return s.roleRepository.Delete(ctx, id)
}
