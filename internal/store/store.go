package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

func NewPGStore(db *sql.DB) Storage {
	return Storage{
		Posts: &PgPostsStore{db: db},
		Users: &PgUsersStore{db: db},
	}
}
