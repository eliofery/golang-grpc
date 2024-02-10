/*
Package eslog ...

Example of using ReplaceAttr.

	slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		    switch a.Key {
		    case slog.SourceKey:
		        return replaceSource(a)
		    case slog.LevelKey:
		        return replaceLevel(a)
		    case slog.TimeKey:
		        return replaceTime(a)
		    default:
		        return a
		    }
		},
	}
*/
package eslog

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// LevelFatal ...
const LevelFatal slog.Level = 12

// LevelNames ...
var LevelNames = map[slog.Leveler]string{
	slog.LevelDebug: "DEB",
	slog.LevelInfo:  "INF",
	slog.LevelWarn:  "WAR",
	slog.LevelError: "ERR",
	LevelFatal:      "FAT",
}

type colorFn func(format string, a ...any) string

// LevelColor ...
var LevelColor = map[slog.Leveler]colorFn{
	slog.LevelDebug: color.HiWhiteString,
	slog.LevelInfo:  color.HiGreenString,
	slog.LevelWarn:  color.HiYellowString,
	slog.LevelError: color.HiMagentaString,
	LevelFatal:      color.HiRedString,
}

// Logger ...
type Logger struct {
	*slog.Logger
	*slog.LevelVar
}

// New log init
// See: https://github.com/golang/go/issues/59145#issuecomment-1481920720
func New(handler Handler) *Logger {
	logger := slog.New(handler)

	return &Logger{
		Logger:   logger,
		LevelVar: handler.LevelVar(),
	}
}

// Log ...
func (l *Logger) log(level slog.Level, msg string, args ...any) {
	if !l.Enabled(context.Background(), level) {
		return
	}

	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)

	_ = l.Handler().Handle(context.Background(), r)
}

// SetLevel ...
func (l *Logger) SetLevel(level slog.Level) *Logger {
	l.LevelVar.Set(level)

	return l
}

// Fatal logs at LevelFatal.
func (l *Logger) Fatal(msg string, args ...any) {
	l.log(LevelFatal, msg, args...)

	os.Exit(1)
}
