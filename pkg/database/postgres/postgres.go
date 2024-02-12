package postgres

import (
	"context"
	"github.com/eliofery/golang-fullstack/pkg/eslog"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres ...
type Postgres struct {
	*pgxpool.Pool
	*config.Config
	*eslog.Logger
}

// New ...
func New(cfg *config.Config, logger *eslog.Logger) *Postgres {
	return &Postgres{
		Config: cfg,
		Logger: logger,
	}
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, p.Config.DSN())
	if err != nil {
		return err
	}
	p.Pool = pool

	if err = p.Pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}
