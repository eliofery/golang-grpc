package config

import (
	"log/slog"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

const (
	logNameDebug = "debug"
	logNameInfo  = "info"
	logNameWarn  = "warn"
	logNameError = "error"
	logNameFatal = "fatal"
)

var (
	// LevelNames ...
	LevelNames = map[string]slog.Level{
		logNameDebug: slog.LevelDebug,
		logNameInfo:  slog.LevelInfo,
		logNameWarn:  slog.LevelWarn,
		logNameError: slog.LevelError,
		logNameFatal: eslog.LevelFatal,
	}

	levelVar *slog.LevelVar
)

// Logger ...
type Logger struct {
	// Level log level
	// debug, info, warn, error, fatal
	Level string `yaml:"level" env-default:"debug"`
}

// LogLevel ...
func (l *Logger) LogLevel() slog.Level {
	level, ok := LevelNames[l.Level]
	if !ok {
		level = LevelNames[logNameInfo]
	}

	return level
}

// LevelVar ...
func (l *Logger) LevelVar() *slog.LevelVar {
	if levelVar == nil {
		levelVar = new(slog.LevelVar)
		levelVar.Set(l.LogLevel())
	}

	return levelVar
}
