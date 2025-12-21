package utils

import "github.com/gin-gonic/gin"

func MustAuth(ctx *gin.Context) AuthContext {
	return ctx.MustGet("auth").(AuthContext)
}
