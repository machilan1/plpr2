package mid

import (
	"context"
	"net/http"

	"github.com/machilan1/plpr2/internal/app/sdk/metrics"
	"github.com/machilan1/plpr2/internal/framework/web"
)

// Metrics updates program counters.
func Metrics() web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx = metrics.Set(ctx)

			err := handler(ctx, w, r)

			n := metrics.AddRequests(ctx)

			if n%1000 == 0 {
				metrics.AddGoroutines(ctx)
			}

			if err != nil {
				metrics.AddErrors(ctx)
			}

			return err
		}

		return h
	}

	return m
}
