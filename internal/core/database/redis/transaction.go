package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type key string

// txKey ...
const txKey key = "txPipeliner"

type transactionManager struct {
	db Transactor
}

// NewTransactionManager ...
func NewTransactionManager(db Client) TxManager {
	return &transactionManager{
		db: db.DB(),
	}
}

// Transaction ...
func (t *transactionManager) Committed(ctx context.Context, fn Handler) ([]redis.Cmder, error) {
	tx, ok := ctx.Value(txKey).(redis.Pipeliner)
	if !ok {
		return nil, fn(ctx)
	}
	defer tx.Discard()

	tx = t.db.TxPipeline()
	ctx = context.WithValue(ctx, txKey, tx)
	if err := fn(ctx); err != nil {
		tx.Discard()

		return nil, err
	}

	return tx.Exec(ctx)
}
