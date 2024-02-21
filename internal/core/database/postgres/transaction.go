package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type key string

// TxKey ...
const TxKey key = "tx"

type transactionManager struct {
	db Transactor
}

// NewTransactionManager ...
func NewTransactionManager(db Client) TxManager {
	return &transactionManager{
		db: db.DB(),
	}
}

// transaction ...
func (t *transactionManager) transaction(ctx context.Context, opts pgx.TxOptions, fn Handler) error {
	tx, ok := ctx.Value(TxKey).(pgx.Tx) // nolint:staticcheck
	if ok {
		return fn(ctx)
	}

	tx, err := t.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	ctx = context.WithValue(ctx, TxKey, tx)
	if err = fn(ctx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %w, rb err: %w", err, rbErr)
		}

		return err
	}

	return tx.Commit(ctx)
}

// Serializable ...
func (t *transactionManager) Serializable(ctx context.Context, cb Handler) error {
	txOpts := pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	}

	return t.transaction(ctx, txOpts, cb)
}

// RepeatableRead ...
func (t *transactionManager) RepeatableRead(ctx context.Context, fn Handler) error {
	txOpts := pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	}

	return t.transaction(ctx, txOpts, fn)
}

// ReadCommitted ...
func (t *transactionManager) ReadCommitted(ctx context.Context, cb Handler) error {
	txOpts := pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}

	return t.transaction(ctx, txOpts, cb)
}

// ReadUncommitted ...
func (t *transactionManager) ReadUncommitted(ctx context.Context, cb Handler) error {
	txOpts := pgx.TxOptions{
		IsoLevel: pgx.ReadUncommitted,
	}

	return t.transaction(ctx, txOpts, cb)
}
