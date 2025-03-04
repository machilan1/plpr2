package healthapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/machilan1/plpr2/internal/framework/web"
)

type handlers struct {
	log *logger.Logger
	db  *sqldb.DB
}

func newHandlers(log *logger.Logger, db *sqldb.DB) *handlers {
	return &handlers{
		log: log,
		db:  db,
	}
}

// readiness checks if the database is ready and if not will return a 500 status.
// Do not respond by just returning an error because further up in the call
// stack it will interpret that as a non-trusted error.
func (h *handlers) readiness(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		h.log.Info(ctx, "readiness failure", "ERROR", err)
		return fmt.Errorf("db not ready: %w", err)
	}

	return web.Respond(ctx, w, AppHealth{Status: "ok"}, http.StatusOK)
}
