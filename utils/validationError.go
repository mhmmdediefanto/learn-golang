package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationError(ctx *gin.Context, err error) {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[strings.ToLower(e.Field())] = e.Translate(Trans)
		}
	}

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Validation error",
		"errors":  errors,
	})
}
