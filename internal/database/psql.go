package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PSQLCreds struct {
	Host, Port, User, Pass, DB string
}

// NewPSQLClient
func NewPSQLClient(creds PSQLCreds) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full", creds.Host, creds.Port, creds.User, creds.Pass, creds.DB)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
