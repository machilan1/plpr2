package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		MaxOpenConns int    `conf:"default:3"`
		DisableTLS   bool   `conf:"default:true"`
		TimeZone     string `conf:"default:Asia/Taipei"`
		// Only set if using CloudSQL. When set, the CloudSQL connector will be used, and `Password`, `Hos`, `DisableTLS`
		// will be ignored.
		// Example: "project:region:instance"
		CloudSQLConnectionName string
	}
	Migration struct {
		Path           string `conf:"default:migrations"`
		TimeoutMinutes int    `conf:"default:10"`
		Down           bool   `conf:"default:false"`
		Version        uint   `conf:"default:0"`
	}
}

func main() {
	log := logger.New(os.Stdout, logger.LevelInfo, "MIGRATE", func(_ context.Context) string {
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

	return dbMigrate(ctx, log, cfg)
}

func dbMigrate(ctx context.Context, log *logger.Logger, cfg config) error {
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

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{
		DatabaseName: cfg.DB.Name,
	})
	if err != nil {
		return fmt.Errorf("failed create migrate driver: %w", err)
	}

	// Run the migrations
	dir := fmt.Sprintf("file://%s", cfg.Migration.Path)
	m, err := migrate.NewWithDatabaseInstance(
		dir,
		cfg.DB.Name,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}
	m.LockTimeout = time.Duration(cfg.Migration.TimeoutMinutes) * time.Minute
	m.Log = newLogger(log)

	if !cfg.Migration.Down {
		log.Info(ctx, "migrate", "msg", "running up migration")
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed run migrate: %w", err)
		}
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			return fmt.Errorf("migrate source error: %w", err)
		}
		if dbErr != nil {
			return fmt.Errorf("migrate database error: %w", err)
		}
	} else {
		log.Info(ctx, "migrate", "msg", "running down migration", "version", cfg.Migration.Version)
		if cfg.Migration.Version == 0 {
			if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				return fmt.Errorf("failed run migrate: %w", err)
			}
		} else {
			if err := m.Migrate(cfg.Migration.Version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				return fmt.Errorf("failed run migrate: %w", err)
			}
		}
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			return fmt.Errorf("migrate source error: %w", err)
		}
		if dbErr != nil {
			return fmt.Errorf("migrate database error: %w", err)
		}
	}

	log.Debug(ctx, "migrate", "msg", "completed successfully")
	return nil
}

type migrateLogger struct {
	log *log.Logger
}

func newLogger(log *logger.Logger) *migrateLogger {
	return &migrateLogger{
		log: logger.NewStdLogger(log, logger.LevelInfo),
	}
}

func (l *migrateLogger) Printf(arg string, vars ...any) {
	l.log.Printf(arg, vars...)
}

func (l *migrateLogger) Verbose() bool {
	return true
}
