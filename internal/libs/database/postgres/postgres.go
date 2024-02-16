package postgres

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres ...
type Postgres struct {
	conn   *pgxpool.Pool
	config *config.Config
	logger *eslog.Logger
}

// New ...
func New(
	ctx context.Context,
	config *config.Config,
	logger *eslog.Logger,
) (*pgxpool.Pool, error) {
	db := &Postgres{
		config: config,
		logger: logger,
	}

	conn, err := db.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if err = db.Migrate(); err != nil {
		return nil, err
	}

	return conn, nil
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context) (*pgxpool.Pool, error) {
	var err error

	p.conn, err = pgxpool.New(ctx, p.config.DSN())
	if err != nil {
		return nil, err
	}

	if err = p.conn.Ping(ctx); err != nil {
		return nil, err
	}

	return p.conn, nil
}
