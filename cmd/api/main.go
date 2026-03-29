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

	chatRepo := repository.NewPostgresChatRepository(database.DB)
	chatHub := services.NewChatHub(chatRepo)
	chatController := controllers.NewChatController(chatHub)

	r := httpx.SetupRouter(userController, chatController)
	r.Run(":8080")
}
