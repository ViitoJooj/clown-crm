package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

var PgUrl string

func GetEnv() {
	godotenv.Load(".env")
	godotenv.Load("../.env")
	godotenv.Load("../../.env")
	godotenv.Load("../../../.env")

	PgUrl = os.Getenv("POSTGRE_URL")
	if PgUrl == "" {
		panic("PgUrl is null")
	}
}
