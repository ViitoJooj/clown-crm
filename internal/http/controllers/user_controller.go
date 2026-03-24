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

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         createdUser.UUID,
		"first_name": createdUser.First_Name,
		"email":      createdUser.Email,
	})
}
