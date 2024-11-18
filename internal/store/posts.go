package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PgPostsStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64    `json:"id"`
	Data      string   `json:"data"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (s *PgPostsStore) Create(ctx context.Context, post *Post) error {

	query := `
	INSERT INTO posts (data, title, user_id, tags) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx, query, post.Data, post.Title, post.UserID, post.Tags, pq.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
