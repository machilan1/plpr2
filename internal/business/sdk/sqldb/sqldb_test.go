package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/machilan1/plpr2/internal/business/sdk/testhelper"
)

// TODO: test nested transaction
// TODO: test nested transaction with different isolation levels

var testDatabaseInstance *TestInstance

func TestMain(m *testing.M) {
	testDatabaseInstance = MustTestInstance()
	defer testDatabaseInstance.MustClose()
	m.Run()
}

func TestDBAfterTransactFails(t *testing.T) {
	t.Parallel()

	log := testhelper.TestLogger(t)
	testDB, _ := testDatabaseInstance.NewDatabase(t, log)

	ctx := context.Background()
	var tx *DB
	err := testDB.Transact(ctx, sql.LevelDefault, func(d *DB) error {
		tx = d
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	var i int
	err = tx.Queryer().QueryRowxContext(ctx, `SELECT 1`).Scan(&i)
	if err == nil {
		t.Fatal("got nil, want error")
	}
}

func TestTransactSerializable(t *testing.T) {
	t.Parallel()

	log := testhelper.TestLogger(t)
	testDB, _ := testDatabaseInstance.NewDatabase(t, log)

	// Test that serializable transactions retry until success.
	// This test was taken from the example at https://www.postgresql.org/docs/11/transaction-iso.html,
	// section 13.2.3.
	ctx := context.Background()

	// Once in a while, the test doesn't work. Repeat to de-flake.
	var msg string
	for i := 0; i < 20; i++ {
		msg = testTransactSerializable(ctx, t, testDB)
		if msg == "" {
			return
		}
	}
	t.Fatal(msg)
}

func testTransactSerializable(ctx context.Context, t *testing.T, db *DB) string {
	const numTransactions = 4
	// A transaction that sums values in class 1 and inserts that sum into class 2,
	// or vice versa.
	insertSum := func(tx *DB, queryClass int) error {
		var sum int
		err := tx.Queryer().QueryRowxContext(ctx, `SELECT SUM(value) FROM ser WHERE class = $1`, queryClass).Scan(&sum)
		if err != nil {
			return err
		}
		insertClass := 3 - queryClass
		_, err = tx.Queryer().ExecContext(ctx, `INSERT INTO ser (class, value) VALUES ($1, $2)`, insertClass, sum)
		return err
	}

	sawRetries := false
	for i := 0; i < 10; i++ {
		for _, stmt := range []string{
			`DROP TABLE IF EXISTS ser`,
			`CREATE TABLE ser (id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY, class INTEGER, value INTEGER)`,
			`INSERT INTO ser (class, value) VALUES (1, 10), (1, 20), (2, 100), (2, 200)`,
		} {
			if _, err := db.Queryer().ExecContext(ctx, stmt); err != nil {
				t.Fatal(err)
			}
		}

		// Run the following two transactions multiple times concurrently:
		//   sum rows with class = 1 and insert as a row with class 2
		//   sum rows with class = 2 and insert as a row with class 1
		errc := make(chan error, numTransactions)
		for i := 0; i < numTransactions; i++ {
			go func() {
				errc <- db.Transact(ctx, sql.LevelSerializable,
					func(tx *DB) error { return insertSum(tx, 1+i%2) })
			}()
		}
		// None of the transactions should fail.
		for i := 0; i < numTransactions; i++ {
			if err := <-errc; err != nil {
				return err.Error()
			}
		}
		t.Logf("max retries: %d", db.MaxRetries())
		// If nothing got retried, this test isn't exercising some important behavior.
		// Try again.
		if db.MaxRetries() > 0 {
			sawRetries = true
			break
		}
	}
	if !sawRetries {
		return "did not see any retries"
	}

	// Demonstrate serializability: there should be numTransactions new rows in
	// addition to the 4 we started with, and viewing the rows in insertion
	// order, each of the new rows should have the sum of the other class's rows
	// so far.
	type row struct {
		Class, Value int
	}
	var rows []row
	if err := QuerySlice(ctx, db, `SELECT class, value FROM ser ORDER BY id`, &rows); err != nil {
		return err.Error()
	}

	const initialRows = 4
	if got, want := len(rows), initialRows+numTransactions; got != want {
		return fmt.Sprintf("got %d rows, want %d", got, want)
	}
	sum := make([]int, 2)
	for i, r := range rows {
		if got, want := r.Value, sum[2-r.Class]; got != want && i >= initialRows {
			return fmt.Sprintf("row #%d: got %d, want %d", i, got, want)
		}
		sum[r.Class-1] += r.Value
	}
	return ""
}
