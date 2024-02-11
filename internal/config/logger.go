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

// LevelNames ...
var LevelNames = map[string]slog.Level{
	logNameDebug: slog.LevelDebug,
	logNameInfo:  slog.LevelInfo,
	logNameWarn:  slog.LevelWarn,
	logNameError: slog.LevelError,
	logNameFatal: eslog.LevelFatal,
}

// Logger ...
type Logger struct {
	// Level log level
	// debug, info, warn, error, fatal
	Level string `yaml:"level" env-default:"debug"`
}

// LogLevel ...
func (c *Config) LogLevel() slog.Level {
	level, ok := LevelNames[c.Logger.Level]
	if !ok {
		level = LevelNames[logNameInfo]
	}

	return level
}
