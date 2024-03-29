package redis

import (
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/redis/go-redis/v9"
)

type clientRedis struct {
	config *Config
	logger eslog.Logger

	masterDBC DB
}

// NewClient ...
func NewClient(config *Config, logger eslog.Logger) Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Address(),
		Username: config.Username(),
		Password: config.Password(),
		DB:       *config.DB(),
	})

	return &clientRedis{
		config: config,
		logger: logger,

		masterDBC: NewDB(conn, logger),
	}
}

// DB ...
func (c *clientRedis) DB() DB {
	return c.masterDBC
}

// Close ...
func (c *clientRedis) Close() error {
	return c.masterDBC.Close()
}
