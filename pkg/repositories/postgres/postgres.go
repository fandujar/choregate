package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func setupPool(ctx context.Context) (*pgxpool.Pool, error) {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		user = "choregate"
	}

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		password = "choregate"
	}

	database := os.Getenv("DATABASE_NAME")
	if database == "" {
		database = "choregate"
	}

	databaseURL := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + database
	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
