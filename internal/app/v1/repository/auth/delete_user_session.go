package auth

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/core/database/redis"

	"github.com/eliofery/golang-grpc/internal/core"
)

// DeleteUserSession ...
func (r *repository) DeleteUserSession(ctx context.Context, userID int64, path string) (int64, error) {
	op := "app.v1.repository.auth.DeleteUserSession"

	q := redis.Query{Name: op}
	keySession := r.redis.KeyFormat(prefixKeySession, userID)
	count, err := r.redis.JSONDel(ctx, q, keySession, path).Result()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	return count, nil
}
