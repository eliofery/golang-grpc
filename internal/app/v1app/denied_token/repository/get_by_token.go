package repository

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// GetByToken ...
func (r *repository) GetByToken(ctx context.Context, token string) (*model.DeniedToken, error) {
	op := "v1.token.repository.GetByToken"

	qb := r.pgQb.
		Select(model.ColumnID, model.ColumnToken).
		From(model.TableName).
		Where(squirrel.Eq{"token": token})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetByToken
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var dToken model.DeniedToken
	if err = r.db.ScanOneContext(ctx, &dToken, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetByToken
	}

	return &dToken, nil
}
