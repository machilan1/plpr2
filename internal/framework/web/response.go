package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/machilan1/plpr2/internal/framework/tracer"
	"go.opentelemetry.io/otel/attribute"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	// If the context has been canceled, it means the client is no longer
	// waiting for a response.
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("client disconnected, do not send response")
		}
	}

	ctx, span := tracer.AddSpan(ctx, "foundation.response", attribute.Int("status", statusCode))
	defer span.End()

	setStatus(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("respond: marshal: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
