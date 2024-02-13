package pretty

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/fatih/color"
)

// AttrFn ...
type AttrFn func(groups []string, attr slog.Attr) slog.Attr

func (h *Handler) dateTime(r slog.Record, rep AttrFn) string {
	timeStr := r.Time.Format(time.DateTime)
	if h.JSON {
		timeStr = r.Time.Format(time.RFC3339Nano)
	}

	if rep != nil {
		if attr := rep(nil, slog.Time(slog.TimeKey, r.Time.Round(0))); attr.Key != "" {
			timeStr = attr.Value.String()
		}
	}

	if h.JSON {
		return timeStr
	}

	return color.WhiteString(timeStr)
}

func (h *Handler) level(r slog.Record, rep AttrFn) string {
	level, ok := eslog.LevelNames[r.Level]
	if !ok {
		level = r.Level.String()
	}

	if rep != nil {
		if attr := rep(nil, slog.Any(slog.LevelKey, r.Level)); attr.Key != "" {
			level = attr.Value.String()
		}
	}

	if h.JSON {
		return level
	}

	return eslog.LevelColor[r.Level](level)
}

func (h *Handler) source(r slog.Record, rep AttrFn) string {
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
		pathSource = fmt.Sprintf("%s:%d", relPath, src.Line)

		if rep != nil {
			if attr := rep(nil, slog.Any(slog.SourceKey, src)); attr.Key != "" {
				pathSource = attr.Value.String()
			}

		}
	}

	return pathSource
}

func (h *Handler) message(r slog.Record) string {
	if h.JSON {
		return r.Message
	}

	return color.HiWhiteString(r.Message)
}

func (h *Handler) attrs(r slog.Record) string {
	if h.JSON {
		var attrsStr []string
		r.Attrs(func(a slog.Attr) bool {
			switch v := a.Value.Any().(type) {
			case string:
				attrsStr = append(attrsStr, fmt.Sprintf("%q:%q", a.Key, v))
			case int64:
				attrsStr = append(attrsStr, fmt.Sprintf("%q:%d", a.Key, v))
			default:
				attrsStr = append(attrsStr, fmt.Sprintf("%q:%v", a.Key, v))
			}

			return true
		})

		return strings.Join(attrsStr, ",")
	}

	attrs := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()

		return true
	})

	b, _ := json.MarshalIndent(attrs, "", "  ")

	return string(b)
}
