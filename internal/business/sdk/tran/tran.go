package tran

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
)

// TxManager controls the lifecycle of a transaction. If any error is returned
// from the transaction function, the transaction will be rolled back.
type TxManager interface {
	RunTx(ctx context.Context, txFunc func(TxManager) error) error
	RunTxWithIsolation(ctx context.Context, iso sql.IsolationLevel, txFunc func(TxManager) error) error
}

// ====================================================================================

// transactor uses the sqldb.DB to manage transactions.
type transactor struct {
	db *sqldb.DB
}

func NewTxManager(db *sqldb.DB) TxManager {
	return &transactor{
		db: db,
	}
}

var _ TxManager = (*transactor)(nil)

func (txM *transactor) RunTx(ctx context.Context, txFunc func(TxManager) error) error {
	return txM.runTx(ctx, sql.LevelDefault, txFunc)
}

func (txM *transactor) RunTxWithIsolation(ctx context.Context, iso sql.IsolationLevel, txFunc func(TxManager) error) error {
	return txM.runTx(ctx, iso, txFunc)
}

func (txM *transactor) runTx(ctx context.Context, iso sql.IsolationLevel, txFunc func(TxManager) error) error {
	fn := func(tx *sqldb.DB) error {
		m := NewTxManager(tx)
		return txFunc(m)
	}

	return txM.db.Transact(ctx, iso, fn)
}

// GetExtContext is a helper function that extracts the underlying sqldb value
// from the TxManager interface for transactional use.
func GetExtContext(tx TxManager) (*sqldb.DB, error) {
	txM, ok := tx.(*transactor)
	if !ok {
		return nil, fmt.Errorf("TxManager(%T) not of type *tran.transactor", tx)
	}

	return txM.db, nil
}
