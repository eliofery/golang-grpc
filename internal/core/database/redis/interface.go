package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Handler ...
type Handler func(ctx context.Context) error

// Query ...
type Query struct {
	Name string
}

// GenericCmdable ...
type GenericCmdable interface {
	KeyFormat(key string, value any) string
	Expire(ctx context.Context, q Query, key string, expiration time.Duration) *redis.BoolCmd
}

// JSONCmdable ...
type JSONCmdable interface {
	JSONSet(ctx context.Context, q Query, key, path string, value any) *redis.StatusCmd
	JSONSetMode(ctx context.Context, q Query, key, path string, value any, mode string) *redis.StatusCmd
	JSONGet(ctx context.Context, q Query, key string, paths ...string) *redis.JSONCmd
	JSONDel(ctx context.Context, q Query, key, path string) *redis.IntCmd
}

// StringCmdable ...
type StringCmdable interface {
	Set(ctx context.Context, q Query, key string, value any, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, q Query, key string) *redis.StringCmd
}

// Cmdabler ...
type Cmdabler interface {
	GenericCmdable
	JSONCmdable
	StringCmdable
}

// Transactor ...
type Transactor interface {
	TxPipeline() redis.Pipeliner
}

// Pinger ...
type Pinger interface {
	Ping(ctx context.Context) error
}

// Closer ...
type Closer interface {
	Close() error
}

// DB ...
type DB interface {
	Cmdabler
	Pinger
	Closer
	Transactor
}

// Client ...
type Client interface {
	DB() DB
	Closer
}

// TxManager ...
type TxManager interface {
	Committed(ctx context.Context, fn Handler) ([]redis.Cmder, error)
}
