package mid

import (
	"context"
	"net/http"

	"github.com/machilan1/plpr2/internal/app/sdk/errs"
	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/machilan1/plpr2/internal/framework/tracer"
	"github.com/machilan1/plpr2/internal/framework/validate"
	"github.com/machilan1/plpr2/internal/framework/web"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {
				log.Error(ctx, "message", "msg", err)

				ctx, span := tracer.AddSpan(ctx, "app.sdk.mid.error")
				span.RecordError(err)
				span.End()

				var er errs.ErrorResponse
				var status int

				switch {
				case errs.IsTrustedError(err):
					trsErr := errs.GetTrustedError(err)

					if validate.IsFieldErrors(trsErr.Err) {
						fieldErrors := validate.GetFieldErrors(trsErr.Err)
						er = errs.ErrorResponse{
							Error:  "data validation error",
							Fields: fieldErrors.Fields(),
						}
						status = trsErr.Status
						break
					}

					er = errs.ErrorResponse{
						Error: trsErr.Error(),
					}
					status = trsErr.Status

				default:
					log.Info(ctx, "unhandled error", "status_code", http.StatusInternalServerError, "error", err)
					er = errs.ErrorResponse{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Respond(ctx, w, er, status); err != nil {
					log.Error(ctx, "web.respond", "error", err)
				}

				// If we receive the shutdown err we need to return it
				// back to the base handler to shut down the service.
				if web.IsShutdown(err) {
					return err
				}
			}

			return nil
		}

		return h
	}

	return m
}
