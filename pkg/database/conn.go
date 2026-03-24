package database

import (
	"context"

	"github.com/ViitoJooj/clown-crm/pkg/dotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Conn() error {
	var err error
	DB, err = pgxpool.New(context.Background(), dotenv.PgUrl)
	if err != nil {
		return err
	}
	return nil
}
