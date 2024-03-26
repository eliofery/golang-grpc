package auth

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core/authorize/token"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
	"github.com/eliofery/golang-grpc/internal/core/metadata"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

const (
	prefixKeySession = "session"
	prefixKeyCache   = "cache:user"

	maxSessions = 5
)

// Repository ...
type Repository interface {
	CreateUserCache(ctx context.Context, user *model.UserCache, expires time.Time) error
	GetUserCache(ctx context.Context, id int64) (*model.UserCache, error)
	DeleteUserCache(ctx context.Context, userID int64, path string) (int64, error)
	CreateUserSession(ctx context.Context, userID int64, expires time.Time) (*model.Token, error)
	GetUserSessions(ctx context.Context, userID int64) (map[string]model.UserSession, error)
	GetUserSession(ctx context.Context, userID int64, key string) (*model.UserSession, error)
	DeleteUserSession(ctx context.Context, userID int64, key string) (int64, error)
}

type repository struct {
	logger          eslog.Logger
	db              postgres.DB
	pgQb            squirrel.StatementBuilderType
	redis           redis.DB
	tokenManager    token.Manager
	metadataManager metadata.Manager
}

// New ...
func New(
	logger eslog.Logger,
	pg postgres.Client,
	redis redis.Client,
	tokenManager token.Manager,
	metadataManager metadata.Manager,
) Repository {
	return &repository{
		logger:          logger,
		db:              pg.DB(),
		pgQb:            pg.QB(),
		redis:           redis.DB(),
		tokenManager:    tokenManager,
		metadataManager: metadataManager,
	}
}
