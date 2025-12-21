package middleware

import (
	"go-bakcend-todo-list/enums"
	"go-bakcend-todo-list/utils"
	"slices"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...enums.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := utils.MustAuth(ctx)
		if slices.Contains(roles, auth.Role) {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(403, gin.H{
			"error": "forbidden",
		})
	}
}
