package utils

import (
	"context"
	"database/sql"
	"fmt"

	ent "github.com/alfariiizi/vandor/internal/infrastructure/db"
)

// TxFunc represents a function that operates within a transaction
type TxFunc func(tx *ent.Tx) error

// WithTx executes a function within a database transaction
// It automatically handles commit/rollback based on the function's return value
func WithTx(ctx context.Context, client *ent.Client, fn TxFunc) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("transaction failed: %v, rollback failed: %w", err, rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// WithTxResult executes a function within a transaction and returns a result
// It automatically handles commit/rollback based on the function's return value
func WithTxResult[T any](ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (T, error)) (T, error) {
	var result T

	tx, err := client.Tx(ctx)
	if err != nil {
		return result, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	result, err = fn(tx)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return result, fmt.Errorf("transaction failed: %v, rollback failed: %w", err, rollbackErr)
		}
		return result, err
	}

	if err := tx.Commit(); err != nil {
		return result, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return result, nil
}

// WithTxOptions executes a function within a transaction with custom options
type TxOptions struct {
	Isolation sql.IsolationLevel
	ReadOnly  bool
}

func WithTxOptions(ctx context.Context, client *ent.Client, opts *TxOptions, fn TxFunc) error {
	var tx *ent.Tx
	var err error

	if opts != nil {
		txOpts := &sql.TxOptions{
			Isolation: opts.Isolation,
			ReadOnly:  opts.ReadOnly,
		}
		tx, err = client.BeginTx(ctx, txOpts)
	} else {
		tx, err = client.Tx(ctx)
	}

	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("transaction failed: %v, rollback failed: %w", err, rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
