package database

import (
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
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
