package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	"github.com/eliofery/golang-fullstack/internal/core/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository ...
type Repository interface {
	Create(context.Context, *model.UserInfo) (*int64, error)
}

type repository struct {
	conn   *pgxpool.Pool
	format squirrel.PlaceholderFormat
}

// New ...
func New(pg *postgres.Postgres) Repository {
	return &repository{
		conn:   pg.DB(),
		format: squirrel.Dollar,
	}
}

// pgQb ...
func (r *repository) pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}
