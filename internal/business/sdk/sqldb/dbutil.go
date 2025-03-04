package sqldb

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/machilan1/plpr2/internal/business/sdk/testhelper"
	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/testcontainers/testcontainers-go"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sethvargo/go-retry"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	// imported to register the postgres migration driver.
	postgresDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	// imported to register the "file" source migration driver.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	// databaseName is the name of the template database to clone.
	databaseName = "test-db-template"

	// databaseUser and databasePassword are the username and password for
	// connecting to the database. These values are only used for testing.
	databaseUser     = "test-user"
	databasePassword = "testing123"

	// defaultPostgresImageRef is the default database container to use if none is
	// specified.
	defaultPostgresImageRef = "postgres:17-alpine"
)

// ApproxTime is a compare helper for clock skew.
var ApproxTime = cmp.Options{cmpopts.EquateApproxTime(1 * time.Second)}

// TestInstance is a wrapper around the Docker-based database instance.
type TestInstance struct {
	log *logger.Logger

	container *postgres.PostgresContainer
	config    Config

	db       *DB
	connLock sync.Mutex

	skipReason string
}

// MustTestInstance is NewTestInstance, except it prints errors to stderr and
// calls os.Exit when finished. Callers can call Close or MustClose().
func MustTestInstance() *TestInstance {
	testDatabaseInstance, err := NewTestInstance()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return testDatabaseInstance
}

// NewTestInstance creates a new Docker-based database instance. It also creates
// an initial database, runs the migrations, and sets that database as a
// template to be cloned by future tests.
//
// This should not be used outside of testing, but it is exposed in the package
// so it can be shared with other packages. It should be called and instantiated
// in TestMain.
//
// All database tests can be skipped by running `go test -short` or by setting
// the `SKIP_DATABASE_TESTS` environment variable.
func NewTestInstance() (*TestInstance, error) {
	// Querying for -short requires flags to be parsed.
	if !flag.Parsed() {
		flag.Parse()
	}

	// Do not create an instance in -short mode.
	if testing.Short() {
		return &TestInstance{
			skipReason: "🚧 Skipping database tests (-short flag provided)!",
		}, nil
	}

	// Do not create an instance if database tests are explicitly skipped.
	if skip, _ := strconv.ParseBool(os.Getenv("SKIP_DATABASE_TESTS")); skip {
		return &TestInstance{
			skipReason: "🚧 Skipping database tests (SKIP_DATABASE_TESTS is set)!",
		}, nil
	}

	ctx := context.Background()

	// Determine the container image to use.
	repository := postgresRepo()

	// Start the actual container.
	container, err := postgres.Run(
		ctx,
		repository,
		postgres.WithDatabase(databaseName),
		postgres.WithUsername(databaseUser),
		postgres.WithPassword(databasePassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start database container: %w", err)
	}

	// TODO: implement this
	// Stop the container after it's been running for too long. Since no test suite
	// should take super long.
	// time.Sleep(120) & container.Terminate(ctx)

	// Build the config
	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}
	port, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}
	hostPort := net.JoinHostPort(host, port.Port())

	cfg := Config{
		User:       databaseUser,
		Password:   databasePassword,
		Host:       hostPort,
		Name:       databaseName,
		DisableTLS: true,
	}

	log := logger.New(io.Discard, logger.LevelInfo, "TEST", func(_ context.Context) string { return "" })

	// Create retryable.
	b := retry.WithMaxRetries(30, retry.NewConstant(1*time.Second))

	// Try to establish a connection to the database, with retries.
	var db *DB
	if err := retry.Do(ctx, b, func(ctx context.Context) error {
		var err error
		db, err = Open(log, cfg)
		if err != nil {
			return retry.RetryableError(err)
		}
		if err := db.Ping(ctx); err != nil {
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed waiting for database container to be ready: %w", err)
	}

	// Run the init scripts.
	if err := dbInit(db); err != nil {
		return nil, fmt.Errorf("failed to run init scripts: %w", err)
	}

	// Run the migrations.
	if err := dbMigrate(cfg); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// Return the instance.
	return &TestInstance{
		log:       log,
		container: container,
		config:    cfg,
		db:        db,
	}, nil
}

// MustClose is like Close except it prints the error to stderr and calls os.Exit.
func (i *TestInstance) MustClose() {
	if err := i.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Close terminates the test database instance, cleaning up any resources.
func (i *TestInstance) Close() (retErr error) {
	// Do not attempt to close  things when there's nothing to close.
	if i.skipReason != "" {
		return
	}

	ctx := context.Background()

	defer func() {
		if err := i.container.Terminate(ctx); err != nil {
			retErr = fmt.Errorf("failed to terminate database container: %w", err)
			return
		}
	}()

	if err := i.db.Close(ctx); err != nil {
		retErr = fmt.Errorf("failed to close connection: %w", err)
		return
	}

	return
}

// NewDatabase creates a new database suitable for use in testing. It returns an
// established database connection and the configuration.
func (i *TestInstance) NewDatabase(tb testing.TB, log *logger.Logger) (*DB, Config) {
	tb.Helper()

	// Ensure we should actually create the database.
	if i.skipReason != "" {
		tb.Skip(i.skipReason)
	}

	// Clone the template database.
	newDatabaseName, err := i.clone()
	if err != nil {
		tb.Fatal(err)
	}

	cfg := i.config
	cfg.Name = newDatabaseName

	// Establish a connection to the database.
	db, err := Open(log, cfg)
	if err != nil {
		tb.Fatalf("failed to connect to database %q: %s", newDatabaseName, err)
	}

	// Close connection and delete database when done.
	tb.Cleanup(func() {
		tb.Helper()

		ctx := context.Background()

		// Close connection first. It is an error to drop a database with active
		// connections.
		db.Close(ctx)

		// Drop the database to keep the container from running out of resources.
		q := fmt.Sprintf(`DROP DATABASE IF EXISTS "%s" WITH (FORCE);`, newDatabaseName)

		i.connLock.Lock()
		defer i.connLock.Unlock()

		if _, err := i.db.db.ExecContext(ctx, q); err != nil {
			tb.Errorf("failed to drop database %q: %s", newDatabaseName, err)
		}
	})

	return db, cfg
}

// dbMigrate runs the migrations. u is the connection URL string (e.g.
// postgres://...).
func dbMigrate(cfg Config) error {
	sqlxDB, err := open(cfg)
	if err != nil {
		return fmt.Errorf("failed to open sqlxDB: %w", err)
	}

	driver, err := postgresDriver.WithInstance(sqlxDB.DB, &postgresDriver.Config{
		DatabaseName: cfg.Name,
	})
	if err != nil {
		return fmt.Errorf("failed create migrate driver: %w", err)
	}

	// Run the migrations
	migrationsDir := fmt.Sprintf("file://%s", dbMigrationsDir())
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, cfg.Name, driver)
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed run migrate: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	return nil
}

func dbInit(db *DB) error {
	initDir := dbInitDir()
	initSQLPath := fmt.Sprintf("%s/init.sql", initDir)
	sql, err := os.ReadFile(initSQLPath)
	if err != nil {
		return fmt.Errorf("readFile: path[%s]: %w", initSQLPath, err)
	}

	if _, err := db.db.Exec(string(sql)); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

// clone creates a new database with a random name from the template instance.
func (i *TestInstance) clone() (string, error) {
	// Generate a random database name.
	name, err := randomDatabaseName()
	if err != nil {
		return "", fmt.Errorf("failed to generate random database name: %w", err)
	}

	// Setup context and create SQL command. Unfortunately we cannot use parameter
	// injection here as that's only valid for prepared statements, for which this
	// is not. Fortunately both inputs can be trusted in this case.
	ctx := context.Background()
	q := fmt.Sprintf(`CREATE DATABASE "%s" WITH TEMPLATE "%s";`, name, databaseName)

	// Unfortunately postgres does not allow parallel database creation from the
	// same template, so this is guarded with a lock.
	i.connLock.Lock()
	defer i.connLock.Unlock()

	// Clone the template database as the new random database name.
	if _, err := i.db.db.ExecContext(ctx, q); err != nil {
		return "", fmt.Errorf("failed to clone template database: %w", err)
	}
	return name, nil
}

// dbMigrationsDir returns the path on disk to the migrations. It uses
// runtime.Caller() to get the path to the caller, since this package is
// imported by multiple others at different levels.
func dbMigrationsDir() string {
	return testhelper.TestDataPath("../../../../migrations")
}

// dbInitDir returns the path on disk to the database initialization scripts.
func dbInitDir() string {
	return testhelper.TestDataPath("../../../../build/postgres")
}

// postgresRepo returns the postgres container image name based on an
// environment variable, or the default value if the environment variable is
// unset.
func postgresRepo() string {
	ref := os.Getenv("CI_POSTGRES_IMAGE")
	if ref == "" {
		ref = defaultPostgresImageRef
	}

	return ref
}

// randomDatabaseName returns a random database name.
func randomDatabaseName() (string, error) {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
