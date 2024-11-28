package db

import (
	"context"
	"database/sql"
	"time"
)

func New(address string, maxOpenConnections, maxIdleConnections int, maxIdleTime string) (*sql.DB, error) {

	db, err := sql.Open("postgres", address)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context is canceled to prevent a resource leak

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
