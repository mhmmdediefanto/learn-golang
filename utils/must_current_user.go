package utils

import "github.com/gin-gonic/gin"

func MustCurrentUser(ctx *gin.Context) uint {
	return ctx.MustGet("userID").(uint)
}
