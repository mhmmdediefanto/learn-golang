package middleware

import (
	"go-bakcend-todo-list/enums"
	"go-bakcend-todo-list/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Implementasi middleware autentikasi di sini
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(ctx, http.StatusUnauthorized, "Authorization Header Missing", nil)
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(ctx, http.StatusUnauthorized, "Invalid Token format", nil)
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, "Invalid or Expired Token", err)
			ctx.Abort()
			return
		}

		// Simpan user ID di context untuk digunakan di handler berikutnya
		ctx.Set("auth", utils.AuthContext{
			UserID: claims.UserID,
			Role:   enums.UserRole(claims.Role),
		})
		ctx.Next()

	}
}
