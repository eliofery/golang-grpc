package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	"github.com/eliofery/golang-fullstack/internal/core/database/postgres"
)

// Repository ...
type Repository interface {
	Create(context.Context, *model.UserInfo) (*int64, error)
}

type repository struct {
	db     postgres.DB
	format squirrel.PlaceholderFormat
}

// New ...
func New(pg postgres.Client) Repository {
	return &repository{
		db:     pg.DB(),
		format: squirrel.Dollar,
	}
}

// pgQb ...
func (r *repository) pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}
