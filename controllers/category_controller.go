package controllers

import (
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetAll(ctx *gin.Context) {
	categories, err := c.service.GetAll()
	if err != nil {
		utils.Error(ctx, 500, "Gagal Mengambil Kategori", err)
		return
	}

	utils.Success(ctx, "Berhasil mengambil kategori", categories)
}
