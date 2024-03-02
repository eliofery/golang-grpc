package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errGetDefaultRole = status.Error(codes.Internal, "role not found")
	errNotFound       = status.Error(codes.NotFound, "setting not found")
)

// Repository ...
type Repository interface {
	GetDefaultRoleID(context.Context) (int64, error)
}

type repository struct {
	db     postgres.DB
	pgQb   squirrel.StatementBuilderType
	logger *eslog.Logger
}

// New ...
func New(pg postgres.Client, logger *eslog.Logger) Repository {
	return &repository{
		db:     pg.DB(),
		pgQb:   pg.QB(),
		logger: logger,
	}
}
