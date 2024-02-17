package logger

import (
	"fmt"
	"log/slog"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/config"
)

const (
	loggerKeyName = "logger"

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

// Config ...
type Config struct {
	Level     string `yaml:"level"`
	AddSource bool   `yaml:"add-source"`
	JSON      bool   `yaml:"json"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(loggerKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("error to populate logger config: %w", err)
	}

	return &conf, nil
}

// Leveler ...
func (c *Config) Leveler() slog.Level {
	level, ok := LevelNames[c.Level]
	if !ok {
		level = LevelNames[logNameInfo]
	}

	return level
}
