package config

import (
	"log/slog"

	"github.com/eliofery/golang-fullstack/internal/cli"
	"github.com/joho/godotenv"
)

// ServerConfig ...
type ServerConfig interface {
	Address() string
}

// DatabaseConfig ...
type DatabaseConfig interface {
	DSN() string
}

// LoggerConfig ...
type LoggerConfig interface {
	GetLevel() slog.Level
}

// Config ...
type Config struct {
	*cli.Options
}

// New ...
func New(cmd *cli.Options) *Config {
	return &Config{Options: cmd}
}

// LoadGoDotEnv ...
func (c *Config) LoadGoDotEnv() error {
	if err := godotenv.Load(c.ConfigPath); err != nil {
		return err
	}

	return nil
}
