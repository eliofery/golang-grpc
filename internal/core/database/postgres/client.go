package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/migrations"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dirMigration = "sql"
)

// clientPostgres ...
type clientPostgres struct {
	config *Config
	logger eslog.Logger

	masterDBC DB
}

// NewClient ...
func NewClient(ctx context.Context, config *Config, logger eslog.Logger) (Client, error) {
	conn, err := pgxpool.New(ctx, config.DSN())
	if err != nil {
		return nil, err
	}

	return &clientPostgres{
		config: config,
		logger: logger,

		masterDBC: NewDB(conn, logger),
	}, nil
}

// DB ...
func (cp *clientPostgres) DB() DB {
	return cp.masterDBC
}

// QB ...
func (cp *clientPostgres) QB() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

// Close ...
func (cp *clientPostgres) Close() error {
	if cp.masterDBC != nil {
		return cp.masterDBC.Close()
	}

	return nil
}

// Migrate ...
func (cp *clientPostgres) Migrate() error {
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
