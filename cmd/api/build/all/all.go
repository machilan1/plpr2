package all

import (
	"github.com/machilan1/plpr2/internal/app/domain/healthapi"
	"github.com/machilan1/plpr2/internal/app/sdk/mux"
	"github.com/machilan1/plpr2/internal/framework/web"
)

func Routes() add { // nolint: revive
	return add{}
}

type add struct{}

func (add) Add(app *web.App, cfg mux.Config) {
	healthapi.Routes(app, healthapi.Config{
		Log: cfg.Log,
		DB:  cfg.DB,
	})
}
