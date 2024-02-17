package userv1

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/eliofery/golang-fullstack/internal/app/postgres"
	"github.com/eliofery/golang-fullstack/internal/repository/user/v1/model"
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

// Get ...
func (r *repository) Get(_ int64) (*model.User, error) {
	return nil, nil
}
