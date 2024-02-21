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

// ClientPostgres ...
type ClientPostgres struct {
	config *Config
	logger *eslog.Logger

	masterDBC DB
}

// NewClient ...
func NewClient(ctx context.Context, config *Config, logger *eslog.Logger) (Client, error) {
	conn, err := pgxpool.New(ctx, config.DSN())
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &ClientPostgres{
		config: config,
		logger: logger,

		masterDBC: NewDB(conn, logger),
	}, nil
}

// DB ...
func (cp *ClientPostgres) DB() DB {
	return cp.masterDBC
}

// Close ...
func (cp *ClientPostgres) Close() error {
	if cp.masterDBC != nil {
		return cp.masterDBC.Close()
	}

	return nil
}

// Migrate ...
func (cp *ClientPostgres) Migrate() error {
	if !cp.config.IsMigration {
		return nil
	}

	goose.SetLogger(cp.logger)
	goose.SetBaseFS(migrations.EmbedMigration)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(cp.masterDBC.Pool())
	if err := goose.Up(db, dirMigration); err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			cp.logger.Error("failed to close database", slog.Any("err", err))
		}
	}(db)

	return nil
}
