package greetapi

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

func (h *handlers) greet(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		h.log.Info(ctx, "readiness failure", "ERROR", err)
		return fmt.Errorf("db not ready: %w", err)
	}

	return web.Respond(ctx, w, "FCKYBTCH", http.StatusOK)
}
