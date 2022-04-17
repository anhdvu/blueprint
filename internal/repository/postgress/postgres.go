package postgres

import (
	"github.com/anhdvu/blueprint/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PGStore represents a CRUD object using Postgres database.
type PGStore struct {
	db *pgxpool.Pool
}

func NewPGStore() *PGStore {
	return &PGStore{}
}

func (pgs PGStore) Create(t *model.Thing) error {
	return nil
}

func (pgs PGStore) Read(id int64) (model.Thing, error) {
	var t model.Thing
	return t, nil
}

func (pgs PGStore) Update(t *model.Thing) error {
	return nil
}

func (pgs PGStore) Delete(id int64) error {
	return nil
}

func (pgs PGStore) All() ([]model.Thing, error) {
	return nil, nil
}
