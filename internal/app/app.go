package app

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

// App ...
type App struct {
	config *config.Config
	logger *eslog.Logger
	conn   *postgres.Postgres
}

// New ...
func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
		logger: eslog.New(cfg.LoggerHandler(), cfg.LoggerLevelVar()),
	}
}

// Run ...
func Run(a *App) error {
	ctx := context.Background()

	if err := a.PostgresInit(ctx); err != nil {
		return err
	}

	return nil
}

// PostgresInit ...
func (a *App) PostgresInit(ctx context.Context) error {
	a.conn = postgres.New(a.config, a.logger)

	if err := a.conn.Connect(ctx); err != nil {
		return err
	}

	if err := a.conn.Migrate(); err != nil {
		return err
	}

	return nil
}
