package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamUint(ctx *gin.Context, key string) (uint, error) {
	value := ctx.Param(key)
	num, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}
