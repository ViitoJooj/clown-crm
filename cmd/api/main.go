package main

import (
	"github.com/ViitoJooj/clown-crm/pkg/database"
	"github.com/ViitoJooj/clown-crm/pkg/dotenv"
)

func main() {
	dotenv.GetEnv()
	database.Conn()
}
