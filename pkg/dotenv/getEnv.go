package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

var PgUrl string
var JwtAccessTokenSecret string
var JwtRefreshTokenSecret string

func GetEnv() {
	godotenv.Load(".env")
	godotenv.Load("../.env")
	godotenv.Load("../../.env")
	godotenv.Load("../../../.env")

	PgUrl = os.Getenv("POSTGRES_URL")
	if PgUrl == "" {
		panic("PgUrl is null")
	}

	JwtAccessTokenSecret = os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	if JwtAccessTokenSecret == "" {
		panic("JwtAccessTokenSecret is null")
	}

	JwtRefreshTokenSecret = os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	if JwtRefreshTokenSecret == "" {
		panic("JwtRefreshTokenSecret is null")
	}
}
