package controllers

import (
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var req dto.SignInDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	users, accessToken, refreshToken, err := c.service.SignIn(req.Email, req.Password)
	if err != nil {
		utils.Error(ctx, 401, "Gagal sign in", err)
		return
	}

	ctx.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", false, false)
	utils.Success(ctx, "Berhasil sign in", gin.H{
		"user":         users,
		"access_token": accessToken,
	})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.Error(ctx, 401, "User ID not found in context", nil)
		return
	}

	newAccessToken, err := utils.GenerateAccessToken(userID.(uint))
	if err != nil {
		utils.Error(ctx, 500, "Failed to generate new access token", err)
		return
	}
	utils.Success(ctx, "Access token refreshed successfully", gin.H{
		"access_token": newAccessToken,
	})
}

func (c *AuthController) Me(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.Error(ctx, 401, "User ID not found in context", nil)
		return
	}
	user, err := c.service.GetUserByID(userID.(uint))
	if err != nil {
		utils.Error(ctx, 404, "User not found", err)
		return
	}

	utils.Success(ctx, "User retrieved successfully", gin.H{
		"user": user,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {

	// cek userID dari context
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.Error(ctx, 401, "User ID not found in context", nil)
		return
	}
	// panggil service logout
	err := c.service.Logout(userID.(uint))
	if err != nil {
		utils.Error(ctx, 500, "Gagal logout", err)
		return
	}
	// Hapus cookie refresh token dengan mengatur nilai kosong dan waktu kedaluwarsa di masa lalu
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, false)
	utils.Success(ctx, "Berhasil logout", nil)
}
