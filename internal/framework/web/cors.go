package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/machilan1/plpr2/internal/framework/tracer"
)

// EnableCORS enables CORS preflight requests to work in the middleware. It
// prevents the MethodNotAllowedHandler from being called. This must be enabled
// for the CORS middleware to work.
func (a *App) EnableCORS(mw MidFunc) {
	a.mw = append(a.mw, mw)

	handler := func(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
		return Respond(ctx, w, "OK", http.StatusOK)
	}
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.StartTrace(r.Context(), a.tracer, "pkg.web.handle", r.RequestURI, w)
		defer span.End()

		ctx = setTraceID(ctx, span.SpanContext().TraceID().String())
		ctx = setValues(ctx, &Values{})

		_ = handler(ctx, w, r)
	}

	finalPath := fmt.Sprintf("%s %s", http.MethodOptions, "/")

	a.mux.HandleFunc(finalPath, h)
}
