package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Handler ...
type Handler func(ctx context.Context) error

// Query ...
type Query struct {
	Name     string
	QueryRaw string
}

// Client ...
type Client interface {
	DB() DB
	QB() squirrel.StatementBuilderType
	Closer
	Migrater
}

// SQLExecer ...
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer ...
type NamedExecer interface {
	GetContext(ctx context.Context, dest any, q Query, args ...any) error
	ScanOneContext(ctx context.Context, dest any, q Query, args ...any) error
	ScanAllContext(ctx context.Context, dest any, q Query, args ...any) error
	ScanRowContext(ctx context.Context, dest any, q Query, args ...any) error
}

// QueryExecer ...
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...any) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...any) pgx.Row
}

// Transactor ...
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Pinger ...
type Pinger interface {
	Ping(ctx context.Context) error
}

// Pooler ...
type Pooler interface {
	Pool() *pgxpool.Pool
}

// Closer ...
type Closer interface {
	Close() error
}

// DB ...
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Pooler
	Closer
}

// Migrater ...
type Migrater interface {
	Migrate() error
}

// TxManager ...
type TxManager interface {
	Serializable(ctx context.Context, f Handler) error
	RepeatableRead(ctx context.Context, f Handler) error
	ReadCommitted(ctx context.Context, f Handler) error
	ReadUncommitted(ctx context.Context, f Handler) error
}
