package config

import "fmt"

// Postgres ...
type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true" env:"POSTGRES_USER"`
	Password string `yaml:"password" env-required:"true" env:"POSTGRES_PASSWORD"`
	Database string `yaml:"database" env-required:"true" env:"POSTGRES_DATABASE"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

// DSN ...
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Database,
		c.Postgres.SSLMode,
	)
}
