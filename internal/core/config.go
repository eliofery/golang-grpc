package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/config"
	"go.uber.org/fx"
)

const (
	appKeyName    = "app"
	appENVDefault = "local"
)

// EnvConfig ...
type EnvConfig struct {
	ENV         string `yaml:"env"`
	IsMigration bool   `yaml:"is-migration"`
}

// Config ...
type Config struct {
	fx.Out

	Provider config.Provider
	Config   *EnvConfig
}

// NewConfig ...
func NewConfig(cli *Options) (Config, error) {
	if _, err := os.Stat(cli.EnvPath); err != nil {
		return Config{}, fmt.Errorf("file %s not found: %s", cli.EnvPath, err)
	}

	if _, err := os.Stat(cli.YamlPath); err != nil {
		return Config{}, fmt.Errorf("file %s not found: %s", cli.YamlPath, err)
	}

	if err := godotenv.Load(cli.EnvPath); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %s", err)
	}

	loader, err := config.NewYAML(config.File(cli.YamlPath))
	if err != nil {
		return Config{}, fmt.Errorf("error loading yaml file: %s", err)
	}

	var conf EnvConfig
	conf.ENV = appENVDefault
	if err = loader.Get(appKeyName).Populate(&conf); err != nil {
		return Config{}, fmt.Errorf("error to populate main config: %w", err)
	}

	return Config{
		Provider: loader,
		Config:   &conf,
	}, nil
}
