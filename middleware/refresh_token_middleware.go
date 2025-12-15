package middleware

import (
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshTokenMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, "Refresh Token Missing", nil)
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateRefreshToken(refreshToken)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, "Invalid or Expired Refresh Token", err)
			ctx.Abort()
			return
		}

		if err := authService.ValidateRefreshToken(claims.UserID, refreshToken); err != nil {
			utils.Error(ctx, http.StatusUnauthorized, err.Error(), nil)
			ctx.Abort()
			return
		}

		// Simpan user ID di context untuk digunakan di handler berikutnya
		ctx.Set("userID", claims.UserID)
		ctx.Set("refresh_token", refreshToken)

		ctx.Next()
	}
}
