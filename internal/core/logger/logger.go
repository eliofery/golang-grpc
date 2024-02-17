package logger

import (
	"log/slog"
	"os"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/eliofery/golang-fullstack/pkg/eslog/pretty"
	"go.uber.org/fx/fxevent"
)

// Logger ...
type Logger struct {
	config *Config
}

// New ...
func New(config *Config) *eslog.Logger {
	logger := Logger{
		config: config,
	}

	return eslog.New(logger.handler(), logger.levelVar())
}

// WithLogger ...
func WithLogger(config *Config) fxevent.Logger {
	return New(config)
}

// Handler ...
func (l *Logger) handler() slog.Handler {
	return pretty.NewHandler(os.Stdout, &pretty.HandlerOptions{
		SlogOptions: &slog.HandlerOptions{
			Level:     l.config.Leveler(),
			AddSource: l.config.AddSource,
		},
		JSON: l.config.JSON,
	})
}

// LevelVar ...
func (l *Logger) levelVar() *slog.LevelVar {
	levelVar := new(slog.LevelVar)
	levelVar.Set(l.config.Leveler())

	return levelVar
}
