package postgres

import (
	"database/sql"
	"log/slog"

	"github.com/eliofery/golang-fullstack/docs"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dirMigration = "migrations"
)

// Migrate ...
func (p *Postgres) Migrate() error {
	if !p.Cli.IsMigration {
		return nil
	}

	goose.SetLogger(p.Logger)
	goose.SetBaseFS(docs.EmbedMigration)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(p.Pool)
	if err := goose.Up(db, dirMigration); err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			p.Logger.Error("failed to close database", slog.Any("err", err))
		}
	}(db)

	return nil
}
