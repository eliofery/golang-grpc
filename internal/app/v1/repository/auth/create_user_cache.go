package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
)

// CreateUserCache ...
func (r *repository) CreateUserCache(ctx context.Context, user *model.UserCache, expires time.Time) error {
	op := "app.v1.repository.auth.CreateUserCache"
	q := redis.Query{Name: op}

	keyCache := r.redis.KeyFormat(prefixKeyCache, user.ID)
	if err := r.redis.JSONSet(ctx, q, keyCache, ".", user).Err(); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrInternal
	}

	if err := r.redis.Expire(ctx, q, keyCache, time.Until(expires)).Err(); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrInternal
	}

	return nil
}
