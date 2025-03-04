package sessdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/machilan1/plpr2/internal/business/sdk/sess"
	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
)

type Store struct {
	db *sqldb.DB
}

func NewStore(db *sqldb.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) QueryByID(ctx context.Context, id string) (sess.Session, error) {
	data := struct {
		SessionID string `db:"session_id"`
	}{
		SessionID: id,
	}

	const q = `
		SELECT session_id,
			   data,
			   created_at,
			   updated_at,
			   expires_at
		FROM sessions
		WHERE session_id = :session_id
		  AND expires_at > CURRENT_TIMESTAMP
	`

	var dbSess dbSession
	if err := sqldb.NamedQueryStruct(ctx, s.db, q, data, &dbSess); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return sess.Session{}, sess.ErrNotFound
		}
		return sess.Session{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toCoreSession(dbSess), nil
}

func (s *Store) Create(ctx context.Context, se sess.Session) error {
	data := toDBSession(se)
	const q = `
		INSERT INTO sessions
			(session_id, data, created_at, updated_at, expires_at)
		VALUES (:session_id, :data, :created_at, :updated_at, :expires_at)
	`
	if err := sqldb.NamedExecContext(ctx, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

func (s *Store) Update(ctx context.Context, se sess.Session) error {
	data := toDBSession(se)
	const q = `
		UPDATE sessions
		SET data       = :data,
			updated_at = :updated_at,
			expires_at = :expires_at
		WHERE session_id = :session_id
	`
	if err := sqldb.NamedExecContext(ctx, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

func (s *Store) Delete(ctx context.Context, id string) error {
	data := struct {
		SessionID string `db:"session_id"`
	}{
		SessionID: id,
	}

	const q = `
		DELETE
		FROM sessions
		WHERE session_id = :session_id
	`
	if err := sqldb.NamedExecContext(ctx, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

func (s *Store) DeleteExpired(ctx context.Context) error {
	const q = `
		DELETE
		FROM sessions
		WHERE expires_at < CURRENT_TIMESTAMP
	`
	if err := sqldb.ExecContext(ctx, s.db, q); err != nil {
		return fmt.Errorf("execcontext: %w", err)
	}

	return nil
}
