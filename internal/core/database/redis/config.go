package redis

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"go.uber.org/config"
)

const configKeyName = "redis"

// Config ...
type Config struct {
	Host    string        `yaml:"host"`
	Port    int           `yaml:"port"`
	Expires time.Duration `yaml:"expires"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate redis config: %w", err)
	}

	return &conf, nil
}

// Address ...
func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// Username ...
func (c *Config) Username() string {
	return os.Getenv("REDIS_USER")
}

// Password ...
func (c *Config) Password() string {
	return os.Getenv("REDIS_PASSWORD")
}

// DB ...
func (c *Config) DB() *int {
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		return nil
	}

	return &db
}
