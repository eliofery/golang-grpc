package auth

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
)

func (r *repository) GetUserCache(ctx context.Context, id int64) (*model.UserCache, error) {
	op := "app.v1.repository.auth.GetUserCache"
	q := redis.Query{Name: op}

	keyCache := r.redis.KeyFormat(prefixKeyCache, id)
	userJSON := r.redis.JSONGet(ctx, q, keyCache, ".").Val()

	var cache model.UserCache
	if err := json.Unmarshal([]byte(userJSON), &cache); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	return &cache, nil
}
