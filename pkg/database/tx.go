package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// transaction abstract
type TxManager interface {
	// Begin(ctx context.Context) (pgx.Tx, error)
	// BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)

	WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error
}

type PgxTxManager struct {
	DB *pgxpool.Pool
}

func NewTxManager(db *pgxpool.Pool) TxManager {
	return &PgxTxManager{
		DB: db,
	}
}

func (m *PgxTxManager) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
