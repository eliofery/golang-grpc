package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres ...
type Postgres struct {
	*pgxpool.Pool
}

// New ...
func New() *Postgres {
	return &Postgres{}
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context, dsn string) error {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}
	p.Pool = pool

	if err = p.Pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}
