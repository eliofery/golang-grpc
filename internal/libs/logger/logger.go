package logger

import (
	"log/slog"
	"os"

	"go.uber.org/fx/fxevent"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/eliofery/golang-fullstack/pkg/eslog/pretty"
)

// Logger ...
type Logger struct {
	config *config.Logger
}

// New ...
func New(config *config.Config) *eslog.Logger {
	var logger Logger
	logger.config = &config.Logger

	handler := logger.handler()
	level := logger.levelVar()

	return eslog.New(handler, level)
}

// WithLogger ...
func WithLogger(config *config.Config) fxevent.Logger {
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
