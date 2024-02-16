package postgres

import (
	"database/sql"
	"log/slog"

	"github.com/eliofery/golang-fullstack/migrations"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dirMigration = "sql"
)

// Migrate ...
func (p *Postgres) Migrate() error {
	if !p.config.Cli.IsMigration {
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
