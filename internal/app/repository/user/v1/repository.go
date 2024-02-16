package userrepositoryv1

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/eliofery/golang-fullstack/internal/app/repository/user/v1/model"
)

// Repository ...
type Repository interface {
	Get(id int64) (*model.User, error)
}

type repository struct {
	conn *pgxpool.Pool
}

// New ...
func New(
	conn *pgxpool.Pool,
) Repository {
	return &repository{
		conn: conn,
	}
}
