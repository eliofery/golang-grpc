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
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"go.uber.org/fx/fxevent"
)

// Level ...
const (
	LevelTrace slog.Level = -8
	LevelFatal slog.Level = 12
)

// LevelNames ...
var LevelNames = map[slog.Leveler]string{
	LevelTrace:      "TRACE",
	slog.LevelDebug: "DEBUG",
	slog.LevelInfo:  "INFO",
	slog.LevelWarn:  "WARN",
	slog.LevelError: "ERROR",
	LevelFatal:      "FATAL",
}

type colorFn func(format string, a ...any) string

// LevelColor ...
var LevelColor = map[slog.Leveler]colorFn{
	LevelTrace:      color.HiWhiteString,
	slog.LevelDebug: color.HiWhiteString,
	slog.LevelInfo:  color.HiGreenString,
	slog.LevelWarn:  color.HiYellowString,
	slog.LevelError: color.HiMagentaString,
	LevelFatal:      color.HiRedString,
}

// LoggerLevel ...
type LoggerLevel interface {
	Trace(msg string, args ...any)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)

	Sprintf(msg string, args ...any) string
	Fatalf(msg string, args ...any)
	Print(msg string, args ...any)
	Printf(msg string, args ...any)
}

// Logger ...
type Logger interface {
	fxevent.Logger
	LoggerLevel
}

// Logger ...
type logger struct {
	*slog.Logger
	*slog.LevelVar
}

// New log init
// See: https://github.com/golang/go/issues/59145#issuecomment-1481920720
func New(handler slog.Handler, lvl *slog.LevelVar) Logger {
	log := slog.New(handler)
	slog.SetDefault(log)

	return &logger{
		Logger:   log,
		LevelVar: lvl,
	}
}

// Log ...
func (l *logger) log(level slog.Level, msg string, args ...any) {
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

// Trace logs at LevelTrace.
func (l *logger) Trace(msg string, args ...any) {
	l.log(LevelTrace, msg, args...)
}

// Fatal logs at LevelFatal.
func (l *logger) Fatal(msg string, args ...any) {
	l.log(LevelFatal, msg, args...)

	os.Exit(1)
}

// Sprintf ...
func (l *logger) Sprintf(msg string, args ...any) string {
	return fmt.Sprintf(msg, args...)
}

// Fatalf logs at LevelFatal.
func (l *logger) Fatalf(msg string, args ...any) {
	l.Fatal(l.Sprintf(l.removeLineBreak(msg), args...))
}

// Print logs any level.
func (l *logger) Print(msg string, args ...any) {
	l.log(l.Level(), l.removeLineBreak(msg), args...)
}

// Printf logs any level.
func (l *logger) Printf(msg string, args ...any) {
	l.log(l.Level(), l.Sprintf(l.removeLineBreak(msg), args...))
}

// removeLineBreak ...
func (l *logger) removeLineBreak(msg string) string {
	return strings.Replace(msg, "\n", " ", -1)
}
