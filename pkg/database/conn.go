package database

import (
	"context"
	"crm/pkg/dotenv"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Conn() {
	DB, err := pgxpool.New(context.Background(), dotenv.PgUrl)
	if err != nil {
		panic(err)
	}

	defer DB.Close()
}
