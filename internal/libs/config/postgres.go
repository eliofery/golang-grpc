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
func (p *Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.Database,
		p.SSLMode,
	)
}
