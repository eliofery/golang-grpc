package postgres

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// postgres ...
type postgres struct {
	conn   *pgxpool.Pool
	logger eslog.Logger
}

// NewDB ...
func NewDB(conn *pgxpool.Pool, logger eslog.Logger) DB {
	return &postgres{
		conn:   conn,
		logger: logger,
	}
}

// Pool ...
func (p *postgres) Pool() *pgxpool.Pool {
	return p.conn
}

// GetContext ...
func (p *postgres) GetContext(ctx context.Context, dest any, q Query, args ...any) error {
	return pgxscan.Get(ctx, p.conn, dest, q.QueryRaw, args...)
}

// ScanOneContext ...
func (p *postgres) ScanOneContext(ctx context.Context, dest any, q Query, args ...any) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAllContext ...
func (p *postgres) ScanAllContext(ctx context.Context, dest any, q Query, args ...any) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// ScanRowContext ...
func (p *postgres) ScanRowContext(ctx context.Context, dest any, q Query, args ...any) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanRow(dest, rows)
}

// ExecContext ...
func (p *postgres) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	p.logger.Debug(q.Name, slog.String("query", q.QueryRaw), slog.Any("args", args))

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.conn.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext ...
func (p *postgres) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	p.logger.Debug(q.Name, slog.String("query", q.QueryRaw), slog.Any("args", args))

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.conn.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext ...
func (p *postgres) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	p.logger.Debug(q.Name, slog.String("query", q.QueryRaw), slog.Any("args", args))

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.conn.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx ...
func (p *postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.conn.BeginTx(ctx, txOptions)
}

// Ping ...
func (p *postgres) Ping(ctx context.Context) error {
	return p.conn.Ping(ctx)
}

// Close ...
func (p *postgres) Close() error {
	p.conn.Close()

	return nil
}
