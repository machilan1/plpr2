package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/machilan1/plpr2/cmd/seed/seeder"
	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
	"github.com/machilan1/plpr2/internal/framework/logger"
)

var build = "develop"

type config struct {
	conf.Version
	DB struct {
		User         string `conf:"default:postgres"`
		Password     string `conf:"default:postgres,mask"`
		Host         string `conf:"default:database-service"`
		Name         string `conf:"default:postgres"`
		MaxIdleConns int    `conf:"default:0"`
		MaxOpenConns int    `conf:"default:0"`
		DisableTLS   bool   `conf:"default:true"`
		TimeZone     string `conf:"default:Asia/Taipei"`
		// Only set if using CloudSQL. When set, the CloudSQL connector will be used, and `Password`, `Hos`, `DisableTLS`
		// will be ignored.
		// Example: "project:region:instance"
		CloudSQLConnectionName string
	}
	Seed struct {
		Version int `conf:"default:0"`
	}
}

func main() {
	log := logger.New(os.Stdout, logger.LevelInfo, "SEED", func(_ context.Context) string {
		return ""
	})

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	cfg := config{
		Version: conf.Version{
			Build: build,
			Desc:  "",
		},
	}

	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	// -----------------------------------------------------------------------------------------------------------------
	// Initialize the database connection.

	log.Info(ctx, "startup", "status", "initializing database", "host", cfg.DB.Host, "name", cfg.DB.Name)

	db, err := sqldb.Open(log, sqldb.Config{
		User:                     cfg.DB.User,
		Password:                 cfg.DB.Password,
		Host:                     cfg.DB.Host,
		Name:                     cfg.DB.Name,
		MaxIdleConns:             cfg.DB.MaxIdleConns,
		MaxOpenConns:             cfg.DB.MaxOpenConns,
		DisableTLS:               cfg.DB.DisableTLS,
		TimeZone:                 cfg.DB.TimeZone,
		CloudSQLInstanceConnName: cfg.DB.CloudSQLConnectionName,
	})
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close(ctx)

	// -----------------------------------------------------------------------------------------------------------------
	// Seed the database.

	log.Info(ctx, "startup", "status", "seeding database", "source", "sql files")

	if err := seeder.SeedSQLFiles(ctx, log, db, cfg.Seed.Version); err != nil {
		return fmt.Errorf("seed sql files: %w", err)
	}

	// -----------------------------------------------------------------------------------------------------------------

	log.Info(ctx, "exit", "status", "completed successfully")

	return nil
}
