package postgres

import (
	"fmt"
	"os"

	"github.com/eliofery/golang-fullstack/internal/core"
	"go.uber.org/config"
)

const postgresKeyName = "postgres"

// Config ...
type Config struct {
	Host        string `yaml:"host"`
	SSLMode     string `yaml:"sslmode"`
	IsMigration bool   `yaml:"is-migration"`
}

// NewConfig ...
func NewConfig(cli *core.Options, provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(postgresKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate postgres config: %w", err)
	}

	if cli.Migration.IsMigration {
		conf.IsMigration = cli.Migration.IsMigration
	}

	return &conf, nil
}

// DSN ...
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		c.SSLMode,
	)
}
