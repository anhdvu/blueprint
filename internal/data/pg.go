package data

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// PGStore represents a CRUD object using Postgres database.
type PGStore struct {
	db *pgxpool.Pool
}

func NewPGStore() *PGStore {
	return &PGStore{}
}

func (pgs PGStore) Create(t *Thing) error {
	return nil
}

func (pgs PGStore) Read(id int64) (*Thing, error) {
	return nil, nil
}

func (pgs PGStore) Update(t *Thing) error {
	return nil
}

func (pgs PGStore) Delete(id int64) error {
	return nil
}
