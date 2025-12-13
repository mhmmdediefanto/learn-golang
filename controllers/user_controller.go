package controllers

import (
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func (c *UserController) GetAll(ctx *gin.Context) {

	users, err := c.userService.GetAll()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal retrieved users", err)
		return
	}

	utils.Success(ctx, "Berhasil retrieved users", users)
}

func (c *UserController) Create(ctx *gin.Context) {
	var userRequest dto.CreateUserDto

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	createdUser, err := c.userService.Create(&user)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal membuat user", err)
		return
	}
	utils.Created(ctx, "User berhasil dibuat", createdUser)
}
