package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

//"postgres://postgres:NeroVit777$@localhost:5432/postgres?sslmode=disable"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
