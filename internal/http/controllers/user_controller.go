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
	var input dtos.InputUserDTO

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

	createdUser, err := c.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := dtos.OutputUserDTO{
		UUID:       createdUser.UUID,
		First_Name: createdUser.First_Name,
		Last_Name:  createdUser.Last_Name,
		Email:      createdUser.Email,
		Updated_at: createdUser.Updated_at.String(),
		Created_at: createdUser.Created_at.String(),
	}

	ctx.JSON(http.StatusCreated, output)
}
