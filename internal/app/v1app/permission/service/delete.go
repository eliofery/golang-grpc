package service

import "context"

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.permissionRepository.Delete(ctx, id)
}
