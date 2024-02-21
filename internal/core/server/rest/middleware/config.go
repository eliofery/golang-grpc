package middleware

import (
	"fmt"

	"go.uber.org/config"
)

const configKeyName = "rest.cors"

// Config ...
type Config struct {
	Origin string `yaml:"origin"`
	MaxAge string `yaml:"max-age"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate rest server config: %w", err)
	}

	return &conf, nil
}
