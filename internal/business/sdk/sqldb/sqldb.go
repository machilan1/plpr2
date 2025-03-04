package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib" // nolint: revive
	"github.com/jmoiron/sqlx"
	"github.com/machilan1/plpr2/internal/framework/logger"
)

// lib/pq errorCodeNames
// https://github.com/lib/pq/blob/master/error.go#L178
const (
	// Class 23 - Integrity Constraint Violation
	integrityConstraintViolation = "23000"
	notNullViolation             = "23502"
	foreignKeyViolation          = "23503"
	uniqueViolation              = "23505"
	checkViolation               = "23514"
	exclusionViolation           = "23P01"
	// Class 40 - Transaction Rollback
	transactionRollback  = "40000"
	serializationFailure = "40001"
	deadlockDetected     = "40P01"
	// Class 42 - Syntax Error or Access Rule Violation
	undefinedTable = "42P01"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound        = sql.ErrNoRows
	ErrDBIntegrity       = errors.New("integrity constraint violation")
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

type DB struct {
	log        *logger.Logger
	db         *sqlx.DB
	tx         *sqlx.Tx
	opts       sql.TxOptions
	mu         sync.Mutex
	maxRetries int
}

// Open creates a new DB for the given configuration.
func Open(log *logger.Logger, cfg Config) (*DB, error) {
	var sqlxDB *sqlx.DB
	var err error
	if cfg.CloudSQLInstanceConnName != "" {
		sqlxDB, err = openCloudSQL(cfg)
	} else {
		sqlxDB, err = open(cfg)
	}
	if err != nil {
		return nil, err
	}

	return New(log, sqlxDB), nil
}

func open(cfg Config) (*sqlx.DB, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	timezone := "utc"
	if cfg.TimeZone != "" {
		timezone = cfg.TimeZone
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", timezone)
	if cfg.Schema != "" {
		q.Set("search_path", cfg.Schema)
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	db, err := sqlx.Open("pgx", u.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func openCloudSQL(cfg Config) (*sqlx.DB, error) {
	tz := "utc"
	if cfg.TimeZone != "" {
		tz = cfg.TimeZone
	}

	d, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithIAMAuthN())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption

	dsn := fmt.Sprintf("user=%s database=%s", cfg.User, cfg.Name)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.ParseConfig: %w", err)
	}
	config.RuntimeParams["timezone"] = tz
	if cfg.Schema != "" {
		config.RuntimeParams["search_path"] = cfg.Schema
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, cfg.CloudSQLInstanceConnName, opts...)
	}
	dbURI := stdlib.RegisterConnConfig(config)

	db, err := sqlx.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}

// New creates a new DB from a sqlx.DB
func New(log *logger.Logger, db *sqlx.DB) *DB {
	return &DB{
		log: log,
		db:  db,
	}
}

// Ping returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func (db *DB) Ping(ctx context.Context) error {
	// If the user doesn't give us a deadline set 1 second.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		if err := db.db.Ping(); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity.
	// Running this query forces a round trip through the database.
	const q = `SELECT TRUE`
	var tmp bool
	return db.db.QueryRowContext(ctx, q).Scan(&tmp)
}

func (db *DB) Close(_ context.Context) error {
	return db.db.Close()
}

// Queryer returns the connection that should be used to execute query
func (db *DB) Queryer() sqlx.ExtContext {
	if db.tx != nil {
		return db.tx
	}
	return db.db
}

// DB returns the underlying *sql.DB.
// This is useful when we need to integrate with third party libraries that only accepts the standard sql.DB type.
func (db *DB) DB() *sql.DB {
	return db.db.DB
}
