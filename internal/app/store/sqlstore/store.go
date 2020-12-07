package sqlstore

import (
	"database/sql"
	"goshop/internal/app/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
	// userRepository *UserRepository
	productRepository *ProductRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}

	return s.productRepository
}
