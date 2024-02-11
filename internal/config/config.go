package config

import (
	"log"

	"github.com/eliofery/golang-fullstack/internal/cli"
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
}

// MustLoad ...
func MustLoad(cmd cli.Options) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(cmd.EnvPath, &cfg); err != nil {
		log.Printf("cannot read env config file: %s", err)
	}

	if err := cleanenv.ReadConfig(cmd.YamlPath, &cfg); err != nil {
		log.Fatalf("cannot read yaml config file: %s", err)
	}

	return &cfg
}
