package mid

import (
	"context"
	"net/http"

	"github.com/machilan1/plpr2/internal/framework/web"
)

// Cors sets the response headers needed for Cross-Origin Resource Sharing
func Cors(origins []string) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			for _, origin := range origins {
				w.Header().Add("Access-Control-Allow-Origin", origin)
			}

			w.Header().Add("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, PATCH, DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
