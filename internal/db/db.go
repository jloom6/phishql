package db

//go:generate retool do mockgen -destination=mocks/db.go -package=mocks github.com/jloom6/phishql/internal/db Interface,Rows

import (
	"context"
	"database/sql"
)

type Interface interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
}

type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

type DB struct {
	sqlDB *sql.DB
}

func New(sqlDB *sql.DB) *DB {
	return &DB{sqlDB: sqlDB}
}

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
