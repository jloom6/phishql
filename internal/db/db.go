package db

//go:generate retool do mockgen -destination=mocks/db.go -package=mocks github.com/jloom6/phishql/internal/db Interface,Rows

import (
	"context"
	"database/sql"
)

// Interface wraps sql.DB's QueryContext method so we can mock it for testing
type Interface interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
}

// Rows wraps sql.Rows's Close, Next, and Scan methods so we can mock it for testing
type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

// DB implements the interface to wrap sql.DB
type DB struct {
	sqlDB *sql.DB
}

// New returns a new sql.DB wrapper
func New(sqlDB *sql.DB) *DB {
	return &DB{sqlDB: sqlDB}
}

// QueryContext call's the underlying sql.DB's QueryContext and wraps the rows to return
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rs, err := db.sqlDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return newRows(rs), nil
}

type rows struct {
	sqlRows *sql.Rows
}

func newRows(rs *sql.Rows) *rows {
	return &rows{sqlRows: rs}
}

func (r *rows) Close() error {
	return r.sqlRows.Close()
}

func (r *rows) Next() bool {
	return r.sqlRows.Next()
}

func (r *rows) Scan(dest ...interface{}) error {
	return r.sqlRows.Scan(dest...)
}
