package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
	// userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
