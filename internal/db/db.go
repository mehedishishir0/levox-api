package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

func Connect(databaseUrl string) (*sql.DB, error) {

	db, err := sql.Open("pgx", databaseUrl)

	if err != nil {
		return nil, fmt.Errorf("sql.Openr error: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("db.PingContext error: %w", err)
	}

	return db, nil

}
