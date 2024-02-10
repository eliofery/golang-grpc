package log

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/fatih/color"
	"github.com/lmittmann/tint"
)

const (
	// LevelFatal ...
	LevelFatal slog.Level = 12
)

// LevelNames ...
// See: https://betterstack.com/community/guides/logging/logging-in-go/
var LevelNames = map[slog.Leveler]string{
	slog.LevelDebug: color.HiWhiteString("DEB"),
	slog.LevelInfo:  color.HiGreenString("INF"),
	slog.LevelWarn:  color.HiYellowString("WAR"),
	slog.LevelError: color.HiMagentaString("ERR"),
	LevelFatal:      color.HiRedString("FAT"),
}

// Logger ...
type Logger struct {
	*slog.Logger
}

// New log init
// See: https://github.com/golang/go/issues/59145#issuecomment-1481920720
func New(conf config.LoggerConfig) *Logger {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      conf.GetLevel(),
		AddSource:  true,
		TimeFormat: "2006/01/02 15:04:05",
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.SourceKey:
				return replaceSource(a)
			case slog.LevelKey:
				return replaceLevel(a)
			default:
				return a
			}
		},
	}))

	return &Logger{
		Logger: logger,
	}
}

// SetLevel ...
func (l *Logger) SetLevel(level slog.Level) *Logger {
	new(slog.LevelVar).Set(level)

	return l
}

// Fatal logs at LevelDebug.
func (l *Logger) Fatal(msg string, args ...any) {
	if !l.Enabled(context.Background(), LevelFatal) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), LevelFatal, msg, pcs[0])
	r.Add(args...)

	_ = l.Handler().Handle(context.Background(), r)

	os.Exit(1)
}
