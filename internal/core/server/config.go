package server

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"go.uber.org/config"
)

const serverKeyName = "server"

// Timeout ...
type Timeout struct {
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
	Idle  time.Duration `yaml:"idle"`
}

// ServGRPC ...
type ServGRPC struct {
	Port int `yaml:"port"`
}

// ServREST ...
type ServREST struct {
	Port int `yaml:"port"`
}

// Config ...
type Config struct {
	Host     string `yaml:"host"`
	Timeout  `yaml:"timeout"`
	ServGRPC `yaml:"grpc"`
	ServREST `yaml:"rest"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get(serverKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("failed to populate server config: %w", err)
	}

	return &conf, nil
}

// GRPCAddress ...
func (c *Config) GRPCAddress() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.ServGRPC.Port))
}

// RESTAddress ...
func (c *Config) RESTAddress() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.ServREST.Port))
}
