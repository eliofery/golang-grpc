package postgres

import (
	"fmt"
	"os"

	"github.com/eliofery/golang-fullstack/internal/app"
	"go.uber.org/config"
)

// Config ...
type Config struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	SSLMode     string `yaml:"sslmode"`
	IsMigration bool   `yaml:"is-migration"`
}

// NewConfig ...
func NewConfig(cli *app.Options, provider config.Provider) (*Config, error) {
	var conf Config
	conf.IsMigration = cli.Migration.IsMigration
	if err := provider.Get("postgres").Populate(&conf); err != nil {
		return nil, fmt.Errorf("postgres config: %w", err)
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
