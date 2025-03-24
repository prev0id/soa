package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), "host=postgres port=5432 user=user password=password dbname=users_db sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return pool, nil
}
