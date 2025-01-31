package transactor

import (
	"context"
	"database/sql"
	"fmt"
)

type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}

type transactor struct {
	db *sql.DB
}

var _ Transactor = &transactor{}

func NewTransactor(db *sql.DB) *transactor {
	return &transactor{db: db}
}

func (t *transactor) WithinTransaction(ctx context.Context, txFunc func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txCtx := txToContext(ctx, tx)
	if err := txFunc(txCtx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

type txCtxKey struct{}

func txToContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey{}, tx)
}
