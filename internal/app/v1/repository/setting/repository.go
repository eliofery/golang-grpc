package setting

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

// Repository ...
type Repository interface {
	GetDefaultRoleByID(context.Context) (int64, error)
}

type repository struct {
	db     postgres.DB
	pgQb   squirrel.StatementBuilderType
	logger eslog.Logger
}

// New ...
func New(pg postgres.Client, logger eslog.Logger) Repository {
	return &repository{
		db:     pg.DB(),
		pgQb:   pg.QB(),
		logger: logger,
	}
}
