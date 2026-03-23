package main

import (
	"crm/pkg/database"
	"crm/pkg/dotenv"
)

func main() {
	dotenv.GetEnv()
	database.Conn()
}
