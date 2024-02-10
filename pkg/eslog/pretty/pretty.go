// Package pretty ...
// See: https://betterstack.com/community/guides/logging/logging-in-go/
package pretty

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

// HandlerOptions ...
type HandlerOptions struct {
	SlogOptions slog.HandlerOptions
	*slog.LevelVar
}

// Handler ...
type Handler struct {
	slog.Handler
	*log.Logger
	HandlerOptions
}

// NewHandler ...
func NewHandler(conf config.LoggerConfig) eslog.Handler {
	lvl := new(slog.LevelVar)
	lvl.Set(conf.GetLevel())

	out := os.Stdout
	opts := HandlerOptions{
		SlogOptions: slog.HandlerOptions{
			Level:     lvl,
			AddSource: true,
		},
		LevelVar: lvl,
	}

	return &Handler{
		Handler:        slog.NewTextHandler(out, &opts.SlogOptions),
		Logger:         log.New(out, "", 0),
		HandlerOptions: opts,
	}
}

// Handle ...
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	rep := h.SlogOptions.ReplaceAttr

	time := h.dateTime(r, rep)
	level := h.level(r, rep)
	source := h.source(r, rep)
	message := h.message(r)
	attrs := h.attrs(r)

	if h.SlogOptions.AddSource {
		h.Logger.Println(time, level, source, message, attrs)
	} else {
		h.Logger.Println(time, level, message, attrs)
	}

	return nil
}

// Enabled ...
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.SlogOptions.Level.Level()
}

// LevelVar ...
func (h *Handler) LevelVar() *slog.LevelVar {
	return h.HandlerOptions.LevelVar
}
