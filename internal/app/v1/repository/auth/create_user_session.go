package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
)

// CreateUserSession ...
func (r *repository) CreateUserSession(ctx context.Context, userID int64, expires time.Time) (*model.Token, error) {
	op := "app.v1.repository.auth.CreateUserSession"

	ipAddress, err := r.metadataManager.GetIPAddress(ctx)
	if err != nil {
		return nil, err
	}

	userAgent, err := r.metadataManager.GetUserAgent(ctx)
	if err != nil {
		return nil, err
	}

	var token model.Token
	token.Refresh, err = r.tokenManager.RefreshToken()
	if err != nil {
		return nil, err
	}

	token.Access, err = r.tokenManager.AccessToken(userID)
	if err != nil {
		return nil, err
	}

	fingerprint := r.tokenManager.GenerateFingerprint(ipAddress, userAgent)

	userSession := model.UserSession{
		UserID:       userID,
		IPAddress:    ipAddress,
		UserAgent:    userAgent,
		RefreshToken: token.Refresh,
		Fingerprint:  fingerprint,
		Expires:      expires,
	}

	userSessions, err := r.GetUserSessions(ctx, userID)
	if err != nil {
		return nil, err
	}
	r.logger.Warn("tst", slog.Any("userSessions", userSessions))

	userSessions[userSession.RefreshToken] = userSession
	if len(userSessions) > maxSessions {
		userSessions = map[string]model.UserSession{
			userSession.RefreshToken: userSession,
		}
	}

	q := redis.Query{Name: op}
	keySession := r.redis.KeyFormat(prefixKeySession, userID)
	if err = r.redis.JSONSet(ctx, q, keySession, ".", userSessions).Err(); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	if err = r.redis.Expire(ctx, q, keySession, time.Until(expires)).Err(); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	return &token, nil
}
