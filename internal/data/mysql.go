package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLStore represents a CRUD object using MySQL database.
type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() *PGStore {
	return &PGStore{}
}

func (mss MySQLStore) Create(t *Thing) error {
	return nil
}

func (mss MySQLStore) Read(id int64) (*Thing, error) {
	return nil, nil
}

func (mss MySQLStore) Update(t *Thing) error {
	return nil
}

func (mss MySQLStore) Delete(id int64) error {
	return nil
}
