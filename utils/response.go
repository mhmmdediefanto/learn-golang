package utils

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(ctx *gin.Context, message string, data any) {
	ctx.JSON(200, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(ctx *gin.Context, message string, data any) {
	ctx.JSON(201, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, status int, message string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	ctx.JSON(status, ApiResponse{
		Success: false,
		Message: message,
		Error:   errMsg,
	})
}
