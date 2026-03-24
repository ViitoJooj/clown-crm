package http

import (
	"github.com/ViitoJooj/clown-crm/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	users := r.Group("/api/v1/users")
	{
		users.POST("/register", userController.Register)
	}

	return r
}
