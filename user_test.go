package crdb_25P02

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lopezator/crdbx"
)

var (
	ctx context.Context
)

func openDB(id string) (*sql.DB, error) {
	db, err := sql.Open("txdb", id)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestMain(m *testing.M) {
	ctx = context.Background()
	txdb.Register("txdb", "crdbx", "postgresql://root@localhost:62527/defaultdb?sslmode=disable")
	os.Exit(m.Run())
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	tests := map[string]struct {
		id string
	}{
		"update1": {id: "01CT3MF9WZN99FA1909BAHXHTN"},
		"update2": {id: "01CT3MJ458P6YDP5E2Q4Z2ZVSP"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if _, err := db.ExecContext(ctx, fmt.Sprintf("UPDATE users SET first_name = 'foo', last_name = 'bar' WHERE id = '%s'", tt.id)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestUpdate2(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	tests := map[string]struct {
		id string
	}{
		"update1": {id: "01CT3MF9WZN99FA1909BAHXHTN"},
		"update2": {id: "01CT3MJ458P6YDP5E2Q4Z2ZVSP"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if _, err := db.ExecContext(ctx, fmt.Sprintf( `UPDATE users SET data = '{"foo": "bar"}' WHERE id = '%s'`, tt.id)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestUpdate3(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	tests := map[string]struct {
		id string
	}{
		"update1": {id: "01CT3MF9WZN99FA1909BAHXHTN"},
		"update2": {id: "01CT3MJ458P6YDP5E2Q4Z2ZVSP"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if _, err := db.ExecContext(ctx, fmt.Sprintf( `UPDATE users SET data = NULL WHERE id = '%s'`, tt.id)); err != nil {
				t.Fatal(err)
			}
		})
	}
}
