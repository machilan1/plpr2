package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/gorilla/sessions"
	"github.com/machilan1/plpr2/cmd/seed/seeder"
	"github.com/machilan1/plpr2/internal/app/sdk/mux"
	"github.com/machilan1/plpr2/internal/business/sdk/blobstore"
	"github.com/machilan1/plpr2/internal/business/sdk/sess"
	"github.com/machilan1/plpr2/internal/business/sdk/sess/stores/sessdb"
	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
	"github.com/machilan1/plpr2/internal/business/sdk/tran"
	"github.com/machilan1/plpr2/internal/framework/logger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var testDatabaseInstance *sqldb.TestInstance

func TestMain(m *testing.M) {
	testDatabaseInstance = sqldb.MustTestInstance()
	defer testDatabaseInstance.MustClose()
	m.Run()
}

func TestServer(t *testing.T) {
	ctx := context.Background()
	log := logger.New(os.Stdout, logger.LevelDebug, "TEST", func(_ context.Context) string { return "" })

	// Enable the development mode of authn middleware.
	// TODO: mid.DevMode = true

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout        time.Duration `conf:"default:5s"`
			WriteTimeout       time.Duration `conf:"default:10s"`
			IdleTimeout        time.Duration `conf:"default:120s"`
			ShutdownTimeout    time.Duration `conf:"default:20s"`
			APIHost            string        `conf:"default:0.0.0.0:3000"`
			DebugHost          string        `conf:"default:0.0.0.0:3010"`
			CORSAllowedOrigins []string      `conf:"default:*"`
		}
		Seed struct {
			Disabled bool `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: build,
		},
	}

	const prefix = ""
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return
		}
		t.Fatalf("parsing config: %v", err)
	}

	// -------------------------------------------------------------------------

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	srv := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      testHandler(t, log, shutdown, cfg.Web.CORSAllowedOrigins, cfg.Seed.Disabled),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		t.Logf("Test server started at %s", cfg.Web.APIHost)
		t.Log("Press Ctrl + C to stop the server")

		serverErrors <- srv.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		t.Fatalf("server error: %v", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "server shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "server shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
			t.Fatalf("could not stop server gracefully: %v", err)
		}
	}
}

func testHandler(
	t *testing.T,
	log *logger.Logger,
	shutdown chan os.Signal,
	corsAllowedOrigins []string,
	seedDisabled bool,
) http.HandlerFunc {
	var srv http.Handler
	return func(w http.ResponseWriter, r *http.Request) {
		if srv == nil {
			// Initial server creation.
			srv = renewServer(t, log, shutdown, corsAllowedOrigins, seedDisabled)
		} else if r.Header.Get("X-Test-Renew") == "True" {
			// Renew the server when requested.
			t.Log("Renewing server...")
			srv = renewServer(t, log, shutdown, corsAllowedOrigins, seedDisabled)
			t.Log("Server renewed")
		}

		srv.ServeHTTP(w, r)
	}
}

func renewServer(t *testing.T, log *logger.Logger, shutdown chan os.Signal, corsAllowedOrigins []string, seedDisabled bool) http.Handler {
	// -------------------------------------------------------------------------
	// Database Support

	testDB, _ := testDatabaseInstance.NewDatabase(t, log)

	if !seedDisabled {
		if err := seeder.SeedSQLFiles(context.Background(), log, testDB, 0); err != nil {
			t.Fatalf("seeding database: %v", err)
		}
	}

	// -------------------------------------------------------------------------
	// Main Session Core Support

	sessOpts := sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30, // 30 days
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	sessCore, err := sess.NewManager(sessdb.NewStore(testDB), sessOpts, []byte("secret-key"))
	if err != nil {
		t.Fatalf("new session core: %v", err)
	}

	// -------------------------------------------------------------------------
	// API Support

	cfgMux := mux.Config{
		Build:    build,
		Shutdown: shutdown,
		Log:      log,
		DB:       testDB,
		Tracer:   sdktrace.NewTracerProvider().Tracer("test tracer"),
		TxM:      tran.NewTxManager(testDB),
		Sess:     sessCore,
		Storage:  blobstore.NewNoopBlobStore(),
	}

	h := mux.WebAPI(cfgMux, buildRoutes(), mux.WithCORS(corsAllowedOrigins))

	return h
}
