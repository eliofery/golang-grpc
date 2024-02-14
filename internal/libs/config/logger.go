package config

import (
	"log/slog"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

const (
	logNameTrace = "trace"
	logNameDebug = "debug"
	logNameInfo  = "info"
	logNameWarn  = "warn"
	logNameError = "error"
	logNameFatal = "fatal"
)

// LevelNames ...
var LevelNames = map[string]slog.Level{
	logNameTrace: eslog.LevelTrace,
	logNameDebug: slog.LevelDebug,
	logNameInfo:  slog.LevelInfo,
	logNameWarn:  slog.LevelWarn,
	logNameError: slog.LevelError,
	logNameFatal: eslog.LevelFatal,
}

// Logger ...
type Logger struct {
	// Level log level
	// trace, debug, info, warn, error, fatal
	Level     string `yaml:"level" env-default:"info"`
	AddSource bool   `yaml:"add-source" env-default:"true"`
	JSON      bool   `yaml:"json" env-default:"false"`
}

// Leveler ...
func (l *Logger) Leveler() slog.Level {
	level, ok := LevelNames[l.Level]
	if !ok {
		level = LevelNames[logNameInfo]
	}

	return level
}
