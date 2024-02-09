package postgres

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres ...
type Postgres struct {
	*pgxpool.Pool
	config.DatabaseConfig
}

// New ...
func New(config config.DatabaseConfig) *Postgres {
	return &Postgres{DatabaseConfig: config}
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, p.DatabaseConfig.DSN())
	if err != nil {
		return err
	}
	p.Pool = pool

	if err = p.Pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}
