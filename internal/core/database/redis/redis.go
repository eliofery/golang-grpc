package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/redis/go-redis/v9"
)

// redis ...
type redisDB struct {
	conn   *redis.Client
	logger eslog.Logger
}

// NewDB ...
func NewDB(conn *redis.Client, logger eslog.Logger) DB {
	return &redisDB{
		conn:   conn,
		logger: logger,
	}
}

// Expire ...
func (r *redisDB) Expire(ctx context.Context, q Query, key string, expiration time.Duration) *redis.BoolCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "Expire"),
		slog.String("key", key),
		slog.Any("expiration", expiration),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.Expire(ctx, key, expiration)
	}

	return r.conn.Expire(ctx, key, expiration)
}

// JSONSet ...
func (r *redisDB) JSONSet(ctx context.Context, q Query, key, path string, value any) *redis.StatusCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "JSONSet"),
		slog.String("path", path),
		slog.String("key", key),
		slog.Any("value", value),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.JSONSet(ctx, key, path, value)
	}

	return r.conn.JSONSet(ctx, key, path, value)
}

// JSONSetMode ...
func (r *redisDB) JSONSetMode(ctx context.Context, q Query, key, path string, value any, mode string) *redis.StatusCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "JSONSetMode"),
		slog.String("path", path),
		slog.String("key", key),
		slog.Any("value", value),
		slog.String("mode", mode),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.JSONSetMode(ctx, key, path, value, mode)
	}

	return r.conn.JSONSetMode(ctx, key, path, value, mode)
}

// JSONGet ...
func (r *redisDB) JSONGet(ctx context.Context, q Query, key string, paths ...string) *redis.JSONCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "JSONGet"),
		slog.Any("paths", paths),
		slog.String("key", key),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.JSONGet(ctx, key, paths...)
	}

	return r.conn.JSONGet(ctx, key, paths...)
}

// JSONDel ...
func (r *redisDB) JSONDel(ctx context.Context, q Query, key, path string) *redis.IntCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "JSONDel"),
		slog.Any("path", path),
		slog.String("key", key),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.JSONDel(ctx, key, path)
	}

	return r.conn.JSONDel(ctx, key, path)
}

// Set ...
func (r *redisDB) Set(ctx context.Context, q Query, key string, value any, expiration time.Duration) *redis.StatusCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "Set"),
		slog.String("key", key),
		slog.Any("value", value),
		slog.Duration("expiration", expiration),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.Set(ctx, key, value, expiration)
	}

	return r.conn.Set(ctx, key, value, expiration)
}

// Get ...
func (r *redisDB) Get(ctx context.Context, q Query, key string) *redis.StringCmd {
	r.logger.Debug(q.Name,
		slog.String("method", "Get"),
		slog.String("key", key),
	)

	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if ok {
		return tx.Get(ctx, key)
	}

	return r.conn.Get(ctx, key)
}

// TxPipeline ...
func (r *redisDB) TxPipeline() redis.Pipeliner {
	return r.conn.TxPipeline()
}

// Ping ...
func (r *redisDB) Ping(ctx context.Context) error {
	return r.conn.Ping(ctx).Err()
}

// Close ...
func (r *redisDB) Close() error {
	return r.conn.Close()
}

// KeyFormat ...
func (r *redisDB) KeyFormat(key string, value any) string {
	return fmt.Sprintf("%s:%v", key, value)
}
