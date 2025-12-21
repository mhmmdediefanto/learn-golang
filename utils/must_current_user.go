package utils

import (
	"go-bakcend-todo-list/enums"

	"github.com/gin-gonic/gin"
)

func MustCurrentUser(ctx *gin.Context) uint {
	return MustAuth(ctx).UserID
}

func MustRoleUser(ctx *gin.Context) enums.UserRole {
	return MustAuth(ctx).Role
}
