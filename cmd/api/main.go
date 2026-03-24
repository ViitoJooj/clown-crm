package main

import (
	httpx "github.com/ViitoJooj/clown-crm/internal/http"
	"github.com/ViitoJooj/clown-crm/internal/http/controllers"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/ViitoJooj/clown-crm/internal/services"
	"github.com/ViitoJooj/clown-crm/pkg/database"
	"github.com/ViitoJooj/clown-crm/pkg/dotenv"
)

func main() {
	dotenv.GetEnv()
	database.Conn()

	userRepo := repository.NewPostgresUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := httpx.SetupRouter(userController)
	r.Run(":8080")
}
