package mysql

import (
	"github.com/jloom6/phishql/internal/db"
)

// Store implements the storage interface for a MySQL database
type Store struct {
	db db.Interface
}

// Params contains the parameters needed to construct a MySQL store
type Params struct {
	DB db.Interface
}

// New constructs a new MySQL store
func New(p Params) *Store {
	return &Store{db: p.DB}
}
