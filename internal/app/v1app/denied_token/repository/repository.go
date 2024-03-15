package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errCreate     = status.Error(codes.Internal, "failed to create token")
	errExists     = status.Error(codes.AlreadyExists, "token already exists")
	errGetByToken = status.Error(codes.Internal, "token not found")
)

// Repository ...
type Repository interface {
	Create(context.Context, string) error
	GetByToken(context.Context, string) (*model.DeniedToken, error)
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
