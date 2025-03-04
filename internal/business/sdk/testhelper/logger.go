package testhelper

import (
	"bytes"
	"context"
	"testing"

	"github.com/machilan1/plpr2/internal/framework/logger"
)

func TestLogger(tb testing.TB) *logger.Logger {
	var buf bytes.Buffer

	tb.Cleanup(func() {
		tb.Helper()

		tb.Log("******************** LOGS ********************\n")
		tb.Logf("\n%s\n", buf.String())
		tb.Log("******************** LOGS ********************\n")
	})

	return logger.New(&buf, logger.LevelDebug, "TEST", func(_ context.Context) string { return "" })
}
