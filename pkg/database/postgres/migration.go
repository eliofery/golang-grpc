package postgres

import (
	"github.com/eliofery/golang-fullstack/internal"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log/slog"
)

const (
	dirMigration = "migrations"
)

// Migrate ...
func (p *Postgres) Migrate() error {
	p.Logger.Info("is migrate", slog.Any("is migrate", p.Cli.IsMigration))
	if !p.Cli.IsMigration {
		return nil
	}

	goose.SetLogger(p.Logger)
	goose.SetBaseFS(internal.EmbedMigration)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(p.Pool)
	if err := goose.Up(db, dirMigration); err != nil {
		return err
	}
	defer db.Close()

	return nil
}
