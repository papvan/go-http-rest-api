package sqlstore

import (
	"database/sql"
	"github.com/papvan/go-http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}