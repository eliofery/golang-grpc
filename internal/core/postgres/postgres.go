package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/eliofery/golang-fullstack/migrations"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dirMigration = "sql"
)

// Postgres ...
type Postgres struct {
	conn   *pgxpool.Pool
	config *Config
	logger *eslog.Logger
}

// New ...
func New(
	ctx context.Context,
	config *Config,
	logger *eslog.Logger,
) (*Postgres, error) {
	pg := &Postgres{
		config: config,
		logger: logger,
	}

	if err := pg.Connect(ctx); err != nil {
		return nil, err
	}

	return pg, nil
}

// Connect ...
func (p *Postgres) Connect(ctx context.Context) error {
	var err error

	p.conn, err = pgxpool.New(ctx, p.config.DSN())
	if err != nil {
		return err
	}

	if err = p.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// DB ...
func (p *Postgres) DB() *pgxpool.Pool {
	return p.conn
}

// Close ...
func (p *Postgres) Close() error {
	p.conn.Close()
	return nil
}

// Ping ...
func (p *Postgres) Ping(ctx context.Context) error {
	return p.conn.Ping(ctx)
}

// Migrate ...
func (p *Postgres) Migrate() error {
	if !p.config.IsMigration {
		return nil
	}

	goose.SetLogger(p.logger)
	goose.SetBaseFS(migrations.EmbedMigration)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(p.conn)
	if err := goose.Up(db, dirMigration); err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			p.logger.Error("failed to close database", slog.Any("err", err))
		}
	}(db)

	return nil
}
