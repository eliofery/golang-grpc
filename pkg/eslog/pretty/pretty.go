// Package pretty ...
// See: https://betterstack.com/community/guides/logging/logging-in-go/
package pretty

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"strings"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

// jsonResult ...
type jsonResult struct {
	Time    string `json:"time,omitempty"`
	Level   string `json:"level,omitempty"`
	Source  string `json:"source,omitempty"`
	Message string `json:"msg,omitempty"`
}

// HandlerOptions ...
type HandlerOptions struct {
	SlogOptions *slog.HandlerOptions
	JSON        bool
}

// Handler ...
type Handler struct {
	slog.Handler
	*log.Logger
	*HandlerOptions
}

// NewHandler ...
func NewHandler(out io.Writer, opts *HandlerOptions) eslog.Handler {
	if opts == nil {
		opts = &HandlerOptions{}
	}

	return &Handler{
		Handler:        slog.NewTextHandler(out, opts.SlogOptions),
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

	var result string
	if h.SlogOptions.AddSource {
		result = fmt.Sprintf("%s %s %s %s %s", time, level, source, message, attrs)
	} else {
		result = fmt.Sprintf("%s %s %s %s", time, level, message, attrs)
	}

	if h.JSON {
		jResult := jsonResult{
			Time:    time,
			Level:   level,
			Source:  source,
			Message: message,
		}

		var jsonData []byte
		jsonData, _ = json.Marshal(jResult)

		// generating json string
		withoutLastChar := string(jsonData[:len(jsonData)-1]) // remove last char "}"
		parts := []string{withoutLastChar}
		if len(attrs) > 0 {
			parts = append(parts, attrs)
		}
		result = strings.Join(parts, ",") + "}" // add last char "}"
	}

	h.Logger.Println(result)

	return nil
}

// Enabled ...
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.SlogOptions.Level.Level()
}
