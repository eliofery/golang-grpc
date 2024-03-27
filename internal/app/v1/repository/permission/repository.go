package permission

import (
	"context"

	"github.com/Masterminds/squirrel"
	permissionModel "github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errNotFound = status.Error(codes.NotFound, "permission not found")
)

// Repository ...
type Repository interface {
	GetByRoleID(ctx context.Context, roleID int64) ([]permissionModel.Permission, error)
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
