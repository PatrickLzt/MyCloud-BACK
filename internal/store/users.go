package store

import (
	"context"
	"database/sql"
)

type PgUsersStore struct {
	db *sql.DB
}

func (s *PgUsersStore) Create(ctx context.Context) error {
	return nil
}
