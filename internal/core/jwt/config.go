package jwt

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/config"
)

const (
	configKeyName = "jwt"

	expiresDefault = 3600
)

// Config ...
type Config struct {
	Secret  string        `yaml:"secret"`
	Expires time.Duration `yaml:"expires"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	conf.Expires = expiresDefault
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("error to populate jwt config: %w", err)
	}

	if conf.Secret == "" {
		return nil, errors.New("jwt secret is required")
	}

	return &conf, nil
}
