package controllers

import (
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/enums"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetAll(ctx *gin.Context) {

	users, err := c.service.GetAll()
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

	if userRequest.Role == "" {
		userRequest.Role = enums.RoleUser
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     userRequest.Role,
	}

	createdUser, err := c.service.Create(&user)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal membuat user", err)
		return
	}
	utils.Created(ctx, "User berhasil dibuat", createdUser)
}
