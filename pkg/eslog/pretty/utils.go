package pretty

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/fatih/color"
)

func (h *Handler) dateTime(r slog.Record, rep eslog.AttrFn) string {
	timeStr := r.Time.Format(time.DateTime)
	if rep != nil {
		if attr := rep(nil, slog.Time(slog.TimeKey, r.Time.Round(0))); attr.Key != "" {
			timeStr = attr.Value.String()
		}
	}

	return color.WhiteString(timeStr)
}

func (h *Handler) level(r slog.Record, rep eslog.AttrFn) string {
	level, ok := eslog.LevelNames[r.Level]
	if !ok {
		level = r.Level.String()
	}

	if rep != nil {
		if attr := rep(nil, slog.Any(slog.LevelKey, r.Level)); attr.Key != "" {
			level = attr.Value.String()
		}
	}

	return eslog.LevelColor[r.Level](level)
}

func (h *Handler) source(r slog.Record, rep eslog.AttrFn) string {
	var pathSource string

	if h.SlogOptions.AddSource {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()

		var src slog.Source
		if f.File != "" {
			src.Function = f.Function
			src.File = f.File
			src.Line = f.Line
		}

		pwd, _ := os.Getwd()
		relPath, _ := filepath.Rel(pwd, src.File)
		basePath := filepath.Base(relPath)
		pathSource = fmt.Sprintf("%s:%d", basePath, src.Line)

		if rep != nil {
			if attr := rep(nil, slog.Any(slog.SourceKey, src)); attr.Key != "" {
				pathSource = attr.Value.String()
			}

		}
	}

	return pathSource
}

func (h *Handler) message(r slog.Record) string {
	return color.CyanString(r.Message)
}

func (h *Handler) attrs(r slog.Record) string {
	attrs := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()

		return true
	})

	b, _ := json.MarshalIndent(attrs, "", "  ")

	return string(b)
}
