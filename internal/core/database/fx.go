package database

import (
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("database",
		fx.Options(
			postgres.NewModule(),
			redis.NewModule(),
		),
	)
}
