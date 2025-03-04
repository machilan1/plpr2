package sessdb

import (
	"time"

	"github.com/machilan1/plpr2/internal/business/sdk/sess"
)

type dbSession struct {
	SessionID string    `db:"session_id"`
	Data      string    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

func toDBSession(s sess.Session) dbSession {
	return dbSession{
		SessionID: s.ID,
		Data:      s.Data,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		ExpiresAt: s.ExpiresAt,
	}
}

func toCoreSession(s dbSession) sess.Session {
	return sess.Session{
		ID:        s.SessionID,
		Data:      s.Data,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		ExpiresAt: s.ExpiresAt,
	}
}
