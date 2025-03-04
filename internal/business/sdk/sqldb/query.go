package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/machilan1/plpr2/internal/framework/tracer"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/attribute"
)

// ExecContext is a helper function to execute a CUD operation with
// logging and tracing.
func ExecContext(ctx context.Context, db *DB, query string) error {
	return NamedExecContext(ctx, db, query, struct{}{})
}

// NamedExecContext is a helper function to execute a CUD operation with
// logging and tracing where field replacement is necessary.
func NamedExecContext(ctx context.Context, db *DB, query string, data any) error {
	return namedExecContext(ctx, db.log, db.Queryer(), query, data, false)
}

// NamedExecContextUsingIn is a helper function to execute a CUD operation with
// logging and tracing where field replacement is necessary. Use this if the
// query has an IN clause.
func NamedExecContextUsingIn(ctx context.Context, db *DB, query string, data any) error {
	return namedExecContext(ctx, db.log, db.Queryer(), query, data, true)
}

func namedExecContext(ctx context.Context, log *logger.Logger, db sqlx.ExtContext, query string, data any, withIn bool) (err error) {
	if data == nil {
		data = struct{}{}
	}

	q := queryString(query, data)

	log.Debug(ctx, "sqldb.NamedExecContext", "query", q)

	defer func() {
		if err != nil {
			switch data.(type) {
			case struct{}:
				log.Infoc(ctx, 6, "sqldb.NamedExecContext", "query", q, "ERROR", err)
			default:
				log.Infoc(ctx, 5, "sqldb.NamedExecContext", "query", q, "ERROR", err)
			}
		}
	}()

	ctx, span := tracer.AddSpan(ctx, "sqldb.NamedExecContext", attribute.String("query", q))
	defer span.End()

	switch withIn {
	case true:
		_, err = func() (sql.Result, error) {
			named, args, err := sqlx.Named(query, data)
			if err != nil {
				return nil, err
			}

			query, args, err := sqlx.In(named, args...)
			if err != nil {
				return nil, err
			}

			query = db.Rebind(query)
			return db.ExecContext(ctx, query, args...)
		}()
	default:
		_, err = sqlx.NamedExecContext(ctx, db, query, data)
	}

	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			switch pgerr.Code {
			case undefinedTable:
				return errors.Join(ErrUndefinedTable, err)
			case uniqueViolation:
				return errors.Join(ErrDBDuplicatedEntry, err)
			case integrityConstraintViolation, notNullViolation, foreignKeyViolation, checkViolation, exclusionViolation:
				return errors.Join(ErrDBIntegrity, err)
			}
		}
		return err
	}

	return nil
}

// QuerySlice is a helper function for executing queries that return a
// collection of data to be unmarshalled into a slice.
func QuerySlice[T any](ctx context.Context, db *DB, query string, dest *[]T) error {
	return namedQuerySlice(ctx, db.log, db.Queryer(), query, struct{}{}, dest, false)
}

// NamedQuerySlice is a helper function for executing queries that return a
// collection of data to be unmarshalled into a slice where field replacement is
// necessary.
func NamedQuerySlice[T any](ctx context.Context, db *DB, query string, data any, dest *[]T) error {
	return namedQuerySlice(ctx, db.log, db.Queryer(), query, data, dest, false)
}

// NamedQuerySliceUsingIn is a helper function for executing queries that return
// a collection of data to be unmarshalled into a slice where field replacement
// is necessary. Use this if the query has an IN clause.
func NamedQuerySliceUsingIn[T any](ctx context.Context, db *DB, query string, data any, dest *[]T) error {
	return namedQuerySlice(ctx, db.log, db.Queryer(), query, data, dest, true)
}

func namedQuerySlice[T any](ctx context.Context, log *logger.Logger, db sqlx.ExtContext, query string, data any, dest *[]T, withIn bool) (err error) {
	if data == nil {
		data = struct{}{}
	}

	q := queryString(query, data)

	log.Debug(ctx, "sqldb.NamedQuerySlice", "query", q)

	defer func() {
		if err != nil {
			log.Infoc(ctx, 6, "sqldb.NamedQuerySlice", "query", q, "ERROR", err)
		}
	}()

	ctx, span := tracer.AddSpan(ctx, "sqldb.NamedQuerySlice", attribute.String("query", q))
	defer span.End()

	var rows *sqlx.Rows

	switch withIn {
	case true:
		rows, err = func() (*sqlx.Rows, error) {
			named, args, err := sqlx.Named(query, data)
			if err != nil {
				return nil, err
			}

			query, args, err := sqlx.In(named, args...)
			if err != nil {
				return nil, err
			}

			query = db.Rebind(query)
			return db.QueryxContext(ctx, query, args...)
		}()
	default:
		rows, err = sqlx.NamedQueryContext(ctx, db, query, data)
	}

	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			switch pgerr.Code {
			case undefinedTable:
				return errors.Join(ErrUndefinedTable, err)
			case uniqueViolation:
				return errors.Join(ErrDBDuplicatedEntry, err)
			case integrityConstraintViolation, notNullViolation, foreignKeyViolation, checkViolation, exclusionViolation:
				return errors.Join(ErrDBIntegrity, err)
			}
		}
		return err
	}
	defer rows.Close()

	var slice []T
	for rows.Next() {
		v := new(T)
		if err := rows.StructScan(v); err != nil {
			return err
		}
		slice = append(slice, *v)
	}
	*dest = slice

	return nil
}

// QueryStruct is a helper function for executing queries that return a
// single value to be unmarshalled into a struct type where field replacement is necessary.
func QueryStruct(ctx context.Context, db *DB, query string, dest any) error {
	return namedQueryStruct(ctx, db.log, db.Queryer(), query, struct{}{}, dest, false)
}

// NamedQueryStruct is a helper function for executing queries that return a
// single value to be unmarshalled into a struct type where field replacement is necessary.
func NamedQueryStruct(ctx context.Context, db *DB, query string, data any, dest any) error {
	return namedQueryStruct(ctx, db.log, db.Queryer(), query, data, dest, false)
}

// NamedQueryStructUsingIn is a helper function for executing queries that return
// a single value to be unmarshalled into a struct type where field replacement
// is necessary. Use this if the query has an IN clause.
func NamedQueryStructUsingIn(ctx context.Context, db *DB, query string, data any, dest any) error {
	return namedQueryStruct(ctx, db.log, db.Queryer(), query, data, dest, true)
}

func namedQueryStruct(ctx context.Context, log *logger.Logger, db sqlx.ExtContext, query string, data any, dest any, withIn bool) (err error) {
	if data == nil {
		data = struct{}{}
	}

	q := queryString(query, data)

	log.Debug(ctx, "sqldb.NamedQueryStruct", "query", q)

	defer func() {
		if err != nil {
			log.Infoc(ctx, 6, "sqldb.NamedQueryStruct", "query", q, "ERROR", err)
		}
	}()

	ctx, span := tracer.AddSpan(ctx, "sqldb.NamedQueryStruct", attribute.String("query", q))
	defer span.End()

	var rows *sqlx.Rows

	switch withIn {
	case true:
		rows, err = func() (*sqlx.Rows, error) {
			named, args, err := sqlx.Named(query, data)
			if err != nil {
				return nil, err
			}

			query, args, err := sqlx.In(named, args...)
			if err != nil {
				return nil, err
			}

			query = db.Rebind(query)
			return db.QueryxContext(ctx, query, args...)
		}()
	default:
		rows, err = sqlx.NamedQueryContext(ctx, db, query, data)
	}

	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			switch pgerr.Code {
			case undefinedTable:
				return errors.Join(ErrUndefinedTable, err)
			case uniqueViolation:
				return errors.Join(ErrDBDuplicatedEntry, err)
			case integrityConstraintViolation, notNullViolation, foreignKeyViolation, checkViolation, exclusionViolation:
				return errors.Join(ErrDBIntegrity, err)
			}
		}
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return ErrDBNotFound
	}

	if err := rows.StructScan(dest); err != nil {
		return err
	}

	return nil
}

// queryString provides a pretty print version of the query and parameters.
// TODO: this implementation won't redact password or other sensitive fields, should be fixed ASAP.
func queryString(query string, args any) string {
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return err.Error()
	}

	for _, param := range params {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("'%s'", v)
		case []byte:
			value = fmt.Sprintf("'%s'", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", value, 1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.Trim(query, " ")
}
