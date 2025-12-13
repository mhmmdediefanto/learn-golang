package middleware

import (
	"go-bakcend-todo-list/repositories"
	"go-bakcend-todo-list/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshTokenMiddleware(userRepo *repositories.UserRepository) gin.HandlerFunc {
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

		dbToken, err := userRepo.GetRefreshTokenByUserID(claims.UserID)
		if err != nil || dbToken != refreshToken {
			utils.Error(ctx, http.StatusUnauthorized, "Refresh Token tidak cocok / sudah di-logout", nil)
			ctx.Abort()
			return
		}

		// Simpan user ID di context untuk digunakan di handler berikutnya
		ctx.Set("userID", claims.UserID)
		ctx.Set("refresh_token", refreshToken)

		ctx.Next()
	}
}
