package logger

import (
	"log/slog"
	"os"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/eliofery/golang-grpc/pkg/eslog/pretty"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Logger ...
type Logger struct {
	fx.Out

	slog.Handler
	*slog.LevelVar
}

// New ...
func New(config *Config) Logger {
	return Logger{
		Handler:  handler(config),
		LevelVar: levelVar(config),
	}
}

// WithLogger ...
func WithLogger(log eslog.Logger) fxevent.Logger {
	return log
}

// Handler ...
func handler(config *Config) slog.Handler {
	return pretty.NewHandler(os.Stdout, &pretty.HandlerOptions{
		SlogOptions: &slog.HandlerOptions{
			Level:     config.Leveler(),
			AddSource: config.AddSource,
		},
		JSON: config.JSON,
	})
}

// LevelVar ...
func levelVar(config *Config) *slog.LevelVar {
	lvl := new(slog.LevelVar)
	lvl.Set(config.Leveler())

	return lvl
}
