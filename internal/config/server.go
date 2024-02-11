package config

import (
	"net"
	"strconv"
	"time"
)

// Timeout ...
type Timeout struct {
	Read  time.Duration `yaml:"read" env-default:"5s"`
	Write time.Duration `yaml:"write" env-default:"5s"`
	Idle  time.Duration `yaml:"idle" env-default:"15s"`
}

// GRPC ...
type GRPC struct {
	Port int `yaml:"port" env-default:"50051"`
}

// REST ...
type REST struct {
	Port int `yaml:"port" env-default:"8080"`
}

// Server ...
type Server struct {
	Host    string `yaml:"host" env-default:"localhost"`
	Timeout `yaml:"timeout"`
	GRPC    `yaml:"grpc"`
	REST    `yaml:"rest"`
}

// GRPCAddress ...
func (c *Config) GRPCAddress() string {
	return net.JoinHostPort(c.Server.Host, strconv.Itoa(c.Server.GRPC.Port))
}

// RESTAddress ...
func (c *Config) RESTAddress() string {
	return net.JoinHostPort(c.Server.Host, strconv.Itoa(c.Server.REST.Port))
}
