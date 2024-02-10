package config

import "log/slog"

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
