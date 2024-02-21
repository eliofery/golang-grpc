package database

import (
	"github.com/eliofery/golang-fullstack/internal/core/database/postgres"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("database",
		fx.Options(
			postgres.NewModule(),
		),
	)
}
