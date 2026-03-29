package http

import (
	"time"

	"github.com/ViitoJooj/clown-crm/internal/http/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, chatController *controllers.ChatController) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	users := r.Group("/api/v1/auth")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.GET("/access-token", userController.AccessToken)
		users.POST("/logout", userController.Logout)

	}

	chat := r.Group("/api/v1/chat")
	{
		chat.GET("/ws", chatController.HandleConnections)
	}

	return r
}
