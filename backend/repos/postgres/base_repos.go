package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseRepository struct {
	pool *pgxpool.Pool
}

// func (r *BaseRepository) BeginTx(ctx context.Context) (context.Context, error) {
// 	tx, err := r.pool.Begin(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return context.WithValue(ctx, "tx", tx), nil
// }

func (r *BaseRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.pool.Begin(ctx)
}

func (r *BaseRepository) CommitTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (r *BaseRepository) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}
