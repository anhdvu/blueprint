package mysql

import (
	"database/sql"

	"github.com/anhdvu/blueprint/internal/model"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLStore represents a CRUD object using MySQL database.
type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() *MySQLStore {
	return &MySQLStore{}
}

func (mss MySQLStore) Create(t *model.Thing) error {
	return nil
}

func (mss MySQLStore) Read(id int64) (model.Thing, error) {
	var t model.Thing
	return t, nil
}

func (mss MySQLStore) Update(t *model.Thing) error {
	return nil
}

func (mss MySQLStore) Delete(id int64) error {
	return nil
}

func (mss MySQLStore) All() ([]model.Thing, error) {
	return nil, nil
}
