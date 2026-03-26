package http

import (
	"github.com/ViitoJooj/clown-crm/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	users := r.Group("/api/v1/auth")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.GET("/access-token", userController.AccessToken)
		users.POST("/logout", userController.Logout)

	}

	return r
}
