package rest

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"go.uber.org/config"
)

const serverKeyName = "rest"

// Timeout ...
type Timeout struct {
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
	Idle  time.Duration `yaml:"idle"`
}

// CORS ...
type CORS struct {
	Origin string `yaml:"origin"`
	MaxAge string `yaml:"max-age"`
}

// Config ...
type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout `yaml:"timeout"`
	CORS    `yaml:"cors"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(serverKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate rest server config: %w", err)
	}

	return &conf, nil
}

// Address ...
func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
