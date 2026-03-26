package controllers

import (
	"net/http"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/http/dtos"
	"github.com/ViitoJooj/clown-crm/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input dtos.InputRegisterDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		First_Name: input.First_Name,
		Last_Name:  input.Last_Name,
		Email:      input.Email,
		Password:   input.Password,
	}

	createdUser, err := c.service.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := dtos.RegisterOutput{
		Success: true,
		Message: "User created.",
		User: dtos.OutputUserDTO{
			UUID:       createdUser.UUID,
			First_Name: createdUser.First_Name,
			Last_Name:  createdUser.Last_Name,
			Email:      createdUser.Email,
			Role:       createdUser.Role,
			Updated_at: createdUser.Updated_at.String(),
			Created_at: createdUser.Created_at.String(),
		},
	}

	ctx.JSON(http.StatusCreated, output)
}

func (c *UserController) Login(ctx *gin.Context) {
	var input dtos.InputLoginDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := dtos.LoginOutput{
		Success: true,
		Message: "Login successful.",
		User: dtos.OutputUserDTO{
			UUID:       user.UUID,
			First_Name: user.First_Name,
			Last_Name:  user.Last_Name,
			Email:      user.Email,
			Role:       user.Role,
			Updated_at: user.Updated_at.String(),
			Created_at: user.Created_at.String(),
		},
		Token: token,
	}

	ctx.JSON(http.StatusOK, output)
}

func (c *UserController) AccessToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token not found",
		})
		return
	}

	_, err = c.service.AccessToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Token is valid",
	})
}

func (c *UserController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successful",
	})
}
