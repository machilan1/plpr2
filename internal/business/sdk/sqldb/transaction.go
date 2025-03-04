package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

func (db *DB) InTransaction() bool {
	return db.tx != nil
}

func (db *DB) IsRetryable() bool {
	return db.tx != nil && isRetryable(db.opts.Isolation)
}

// Transact executes the given function in the context of a SQL transaction at
// the given isolation level, rolling back the transaction if the function
// panics or returns an error.
//
// The given function is called with a DB that is associated with a transaction.
// The DB should be used only inside the function; if it is used to access the
// database after the function returns, the calls will return errors.
//
// If the isolation level requires it, Transact will retry the transaction upon
// serialization failure, so txFunc may be called more than once.
func (db *DB) Transact(ctx context.Context, iso sql.IsolationLevel, txFunc func(db *DB) error) error {
	opts := &sql.TxOptions{Isolation: iso}
	if isRetryable(iso) {
		return db.transactWithRetry(ctx, opts, txFunc)
	}

	return db.transact(ctx, opts, txFunc)
}

func (db *DB) transactWithRetry(ctx context.Context, opts *sql.TxOptions, txFunc func(*DB) error) (err error) {
	// Retry on serialization failure, up to some max.
	const maxRetries = 10
	sleepDur := 125 * time.Millisecond
	for i := 0; i <= maxRetries; i++ {
		err = db.transact(ctx, opts, txFunc)
		if isSerializationFailure(err) {
			db.mu.Lock()
			if i > db.maxRetries {
				db.maxRetries = i
			}
			db.mu.Unlock()
			db.log.Debug(ctx, fmt.Sprintf("serialization failure; retrying after %s", sleepDur))
			time.Sleep(sleepDur)
			sleepDur *= 2
			continue
		}
		if err != nil {
			db.log.Debug(ctx, "transactWithRetry", "error", err)
			if strings.Contains(err.Error(), serializationFailure) {
				return fmt.Errorf("error text has %q but not recognized as serialization failure: type %T, err %w", serializationFailure, err, err) // nolint:errorlint
			}
		}
		if i > 0 {
			db.log.Debug(ctx, fmt.Sprintf("retried serializable transaction %d time(s)", i))
		}
		return err
	}
	return fmt.Errorf("reached max number of tries due to serialization failure (%d)", maxRetries)
}

func isSerializationFailure(err error) bool {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		return pgerr.Code == serializationFailure
	}
	return false
}

func (db *DB) transact(ctx context.Context, opts *sql.TxOptions, txFunc func(*DB) error) (err error) {
	if db.InTransaction() {
		return errors.New("already in transaction")
	}

	tx, err := db.db.BeginTxx(ctx, opts)
	if err != nil {
		return fmt.Errorf("db.BeginTxx(): %w", err)
	}

	db.log.Debug(ctx, "transaction started", "isolation", opts.Isolation.String())
	start := time.Now()
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			if txErr := tx.Commit(); txErr != nil {
				// To propagate this error correctly, we must use a named return values
				err = fmt.Errorf("tx.Commit(): %w", txErr)
			}
		}

		db.log.Debug(ctx, "transaction finished",
			"isolation", opts.Isolation.String(),
			"duration", float64(time.Since(start))/1000000.0, // milliseconds
			"error", err,
		)
	}()

	dbtx := New(db.log, db.db)
	dbtx.tx = tx
	dbtx.opts = *opts

	if err := txFunc(dbtx); err != nil {
		return err
	}
	return nil
}

// MaxRetries returns the maximum number of times that a serializable transaction was retried.
func (db *DB) MaxRetries() int {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.maxRetries
}

func isRetryable(iso sql.IsolationLevel) bool {
	return iso == sql.LevelRepeatableRead || iso == sql.LevelSerializable
}
