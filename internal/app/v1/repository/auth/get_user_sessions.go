package auth

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
)

// GetUserSessions ...
func (r *repository) GetUserSessions(ctx context.Context, userID int64) (map[string]model.UserSession, error) {
	op := "app.v1.repository.auth.GetUserSessions"

	q := redis.Query{Name: op}
	keySession := r.redis.KeyFormat(prefixKeySession, userID)
	sessionsData := r.redis.JSONGet(ctx, q, keySession).Val()

	sessions := make(map[string]model.UserSession, maxSessions)
	if sessionsData != "" {
		if err := json.Unmarshal([]byte(sessionsData), &sessions); err != nil {
			r.logger.Debug(op, slog.String("err", err.Error()))
			return nil, core.ErrInternal
		}
	}

	return sessions, nil
}
