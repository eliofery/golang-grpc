package env

import (
	"log/slog"
	"os"

	"github.com/eliofery/golang-fullstack/internal/config"
)

const (
	environmentEnv     = "ENV"
	environmentLocal   = "local"
	environmentProduct = "prod"
)

var levels = map[string]slog.Level{
	environmentLocal:   slog.LevelDebug,
	environmentProduct: slog.LevelInfo,
}

type loggerConfig struct {
	environment string
}

// NewLoggerConfig ...
func NewLoggerConfig() config.LoggerConfig {
	env := os.Getenv(environmentEnv)
	if len(env) == 0 {
		env = environmentLocal
	}

	return &loggerConfig{
		environment: env,
	}
}

// GetLevel level log
func (conf *loggerConfig) GetLevel() slog.Level {
	level, ok := levels[conf.environment]
	if !ok {
		level = levels[environmentLocal]
	}

	return level
}
