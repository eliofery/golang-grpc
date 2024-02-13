package config

import (
	"log/slog"
	"os"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/eliofery/golang-fullstack/pkg/eslog/pretty"
)

const (
	logNameTrace = "trace"
	logNameDebug = "debug"
	logNameInfo  = "info"
	logNameWarn  = "warn"
	logNameError = "error"
	logNameFatal = "fatal"
)

var (
	// LevelNames ...
	LevelNames = map[string]slog.Level{
		logNameTrace: eslog.LevelTrace,
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
	// trace, debug, info, warn, error, fatal
	Level     string `yaml:"level" env-default:"debug"`
	AddSource bool   `yaml:"add-source" env-default:"true"`
	JSON      bool   `yaml:"json" env-default:"false"`
}

// LoggerHandler ...
func (c *Config) LoggerHandler() slog.Handler {
	return pretty.NewHandler(os.Stdout, &pretty.HandlerOptions{
		SlogOptions: &slog.HandlerOptions{
			Level:     c.LoggerLevel(),
			AddSource: c.AddSource,
		},
		JSON: c.JSON,
	})
}

// LoggerLevel ...
func (c *Config) LoggerLevel() slog.Level {
	level, ok := LevelNames[c.Level]
	if !ok {
		level = LevelNames[logNameInfo]
	}

	return level
}

// LoggerLevelVar ...
func (c *Config) LoggerLevelVar() *slog.LevelVar {
	if levelVar == nil {
		levelVar = new(slog.LevelVar)
		levelVar.Set(c.LoggerLevel())
	}

	return levelVar
}
