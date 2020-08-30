package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type PSQLCreds struct {
	Host, Port, User, Pass, DB string
}

// NewPSQLConn
func NewPSQLConn(ctx context.Context, creds PSQLCreds) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full", creds.User, creds.Pass, creds.Host, creds.Port, creds.DB)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
