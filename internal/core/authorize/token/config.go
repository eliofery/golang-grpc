package token

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/config"
)

const (
	configKeyName = "jwt"

	expiresAccessTokenDefault  = 900    // 15m
	expiresRefreshTokenDefault = 604800 // 7d
)

// Expires ...
type Expires struct {
	AccessToken  time.Duration `yaml:"access-token"`
	RefreshToken time.Duration `yaml:"refresh-token"`
}

// Config ...
type Config struct {
	Secret  string `yaml:"secret"`
	Expires `yaml:"expires"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	conf.Expires.AccessToken = expiresAccessTokenDefault
	conf.Expires.RefreshToken = expiresRefreshTokenDefault
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("error to populate jwt config: %w", err)
	}

	if conf.Secret == "" {
		return nil, errors.New("jwt secret is required")
	}

	return &conf, nil
}
