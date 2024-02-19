package grpc

import (
	"fmt"
	"net"
	"strconv"

	"go.uber.org/config"
)

const serverKeyName = "grpc"

// Config ...
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(serverKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate grpc server config: %w", err)
	}

	return &conf, nil
}

// Address ...
func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
