// Package env ...
package env

import (
	"errors"
	"net"
	"os"

	"github.com/eliofery/golang-fullstack/internal/config"
)

const (
	restHostEnv = "REST_HOST"
	restPortEnv = "REST_PORT"
)

type restConfig struct {
	host string
	port string
}

// NewRESTConfig ...
func NewRESTConfig() (config.ServerConfig, error) {
	host := os.Getenv(restHostEnv)
	if len(host) == 0 {
		return nil, errors.New("rest host not found")
	}

	port := os.Getenv(restPortEnv)
	if len(port) == 0 {
		return nil, errors.New("rest port not found")
	}

	return &restConfig{
		host: host,
		port: port,
	}, nil
}

// Address ...
func (cfg *restConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
