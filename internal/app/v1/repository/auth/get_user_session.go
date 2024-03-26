package auth

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
)

// GetUserSession ...
func (r *repository) GetUserSession(ctx context.Context, userID int64, key string) (*model.UserSession, error) {
	op := "app.v1.repository.auth.GetUserSession"

	q := redis.Query{Name: op}
	keySession := r.redis.KeyFormat(prefixKeySession, userID)
	sessionsData := r.redis.JSONGet(ctx, q, keySession, key).Val()

	var session model.UserSession
	if sessionsData != "" {
		if err := json.Unmarshal([]byte(sessionsData), &session); err != nil {
			r.logger.Debug(op, slog.String("err", err.Error()))
			return nil, core.ErrInternal
		}
	}

	return &session, nil
}
