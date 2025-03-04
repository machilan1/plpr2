package greetapi

import (
	"net/http"

	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/machilan1/plpr2/internal/framework/web"
)

type Config struct {
	Log *logger.Logger
	DB  *sqldb.DB
}

func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newHandlers(cfg.Log, cfg.DB)

	app.HandleFuncNoMid(http.MethodGet, version, "/greet", api.greet)
}
