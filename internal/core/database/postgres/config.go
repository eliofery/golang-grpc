package postgres

import (
	"fmt"
	"os"

	"github.com/eliofery/golang-grpc/internal/core"
	"go.uber.org/config"
)

const configKeyName = "postgres"

// Config ...
type Config struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	SSLMode     string `yaml:"sslmode"`
	IsMigration bool   `yaml:"is-migration"`
}

// NewConfig ...
func NewConfig(cli *core.Options, provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate postgres config: %w", err)
	}

	if cli.Migration.IsMigration {
		conf.IsMigration = cli.Migration.IsMigration
	}

	return &conf, nil
}

// DSN ...
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		c.SSLMode,
	)
}
