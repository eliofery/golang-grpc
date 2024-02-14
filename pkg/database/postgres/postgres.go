package postgres

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/database"
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
func New(config *config.Config, logger *eslog.Logger) database.Database {
	return &Postgres{
		config: config,
		logger: logger,
	}
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context) error {
	var err error

	p.conn, err = pgxpool.New(ctx, p.config.DSN())
	if err != nil {
		return err
	}

	if err = p.conn.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// Close ...
func (p *Postgres) Close() {
	p.conn.Close()
}
