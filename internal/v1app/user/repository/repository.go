package repository

import (
	"github.com/eliofery/golang-fullstack/internal/core/postgres"
	"github.com/eliofery/golang-fullstack/internal/v1app/user/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository ...
type Repository interface {
	Get(id int64) (*model.User, error)
}

type repository struct {
	conn *pgxpool.Pool
}

// New ...
func New(pg *postgres.Postgres) Repository {
	return &repository{
		conn: pg.DB(),
	}
}
