package web

import (
	"context"
)

type ctxKey int

const (
	traceKey ctxKey = iota + 1
	valuesKey
)

// Values represent state for each request.
type Values struct {
	StatusCode int
}

func setTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceKey, traceID)
}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(traceKey).(string)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}

	return v
}

func setStatus(ctx context.Context, status int) {
	v, ok := ctx.Value(valuesKey).(*Values)
	if !ok {
		return
	}

	v.StatusCode = status
}

func setValues(ctx context.Context, v *Values) context.Context {
	return context.WithValue(ctx, valuesKey, v)
}

// GetValues returns the values from the context.
func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(valuesKey).(*Values)
	if !ok {
		return &Values{
			StatusCode: 0,
		}
	}

	return v
}
