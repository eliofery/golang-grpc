// Package env ...
package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/eliofery/golang-fullstack/internal/config"
)

const (
	postgresHostEnv     = "POSTGRES_HOST"
	postgresPortEnv     = "POSTGRES_PORT"
	postgresUserEnv     = "POSTGRES_USER"
	postgresPasswordEnv = "POSTGRES_PASSWORD" //nolint:gosec
	postgresDatabaseEnv = "POSTGRES_DATABASE"
	postgresSSLModeEnv  = "POSTGRES_SSLMODE"

	postgresHostDefault    = "localhost"
	postgresPortDefault    = 5432
	postgresSSLModeDefault = "disable"
)

type postgresConfig struct {
	host     string
	port     int
	user     string
	password string
	database string
	sslMode  string
}

// NewPostgresConfig ...
func NewPostgresConfig() (config.DatabaseConfig, error) {
	host := os.Getenv(postgresHostEnv)
	if len(host) == 0 {
		host = postgresHostDefault
	}

	port, err := strconv.Atoi(os.Getenv(postgresPortEnv))
	if err != nil {
		port = postgresPortDefault
	}

	user := os.Getenv(postgresUserEnv)
	if len(user) == 0 {
		return nil, errors.New("postgres user not found")
	}

	password := os.Getenv(postgresPasswordEnv)
	if len(password) == 0 {
		return nil, errors.New("postgres password not found")
	}

	database := os.Getenv(postgresDatabaseEnv)
	if len(database) == 0 {
		return nil, errors.New("postgres database not found")
	}

	sslMode := os.Getenv(postgresSSLModeEnv)
	if len(sslMode) == 0 {
		sslMode = postgresSSLModeDefault
	}

	return &postgresConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
		sslMode:  sslMode,
	}, nil
}

// DSN ...
func (cfg *postgresConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.user,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.database,
		cfg.sslMode,
	)
}
