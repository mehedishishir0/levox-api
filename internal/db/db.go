package db

import (
	"database/sql"
	"fmt"
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

	return db, nil

}
