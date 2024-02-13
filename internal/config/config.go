package config

import (
	"github.com/eliofery/golang-fullstack/docs/cli"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config ...
type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	Swagger  `yaml:"swagger"`
	Postgres `yaml:"postgres"`
	Adminer  `yaml:"adminer"`

	Cli *cli.Options
}

// New ...
func New(cmd *cli.Options) (*Config, error) {
	var cfg Config
	cfg.Cli = cmd

	if err := cleanenv.ReadConfig(cmd.EnvPath, &cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(cmd.YamlPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
