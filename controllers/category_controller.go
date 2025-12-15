package controllers

import (
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/models"
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

func (c *CategoryController) Create(ctx *gin.Context) {
	var categoryDto dto.CreateCategoryDto
	if err := ctx.ShouldBindJSON(&categoryDto); err != nil {
		utils.ValidationError(ctx, err)
		return
	}
	category := &models.Category{
		Title: categoryDto.Title,
	}
	createdCategory, err := c.service.Create(category)
	if err != nil {
		utils.Error(ctx, 500, "Gagal membuat kategori", err)
		return
	}
	utils.Success(ctx, "Berhasil membuat kategori", createdCategory)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	// Implement delete functionality here
	id, err := utils.ParamUint(ctx, "id")
	if err != nil {
		utils.Error(ctx, 400, "Invalid category ID", err)
		return
	}

	if err := c.service.Delete(id); err != nil {
		utils.Error(ctx, 500, "Gagal menghapus kategori", err)
	}

	utils.Success(ctx, "Berhasil menghapus kategori", nil)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	id, err := utils.ParamUint(ctx, "id")
	if err != nil {
		utils.Error(ctx, 400, "Invalid category ID", err)
		return
	}

	var categoryDto dto.UpdateCategoryDto
	if err := ctx.ShouldBindJSON(&categoryDto); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	dataUpdate := &models.Category{
		Title: categoryDto.Title,
	}

	data, err := c.service.Update(id, dataUpdate)
	if err != nil {
		utils.Error(ctx, 500, "Gagal memperbarui kategori", err)
		return
	}
	utils.Success(ctx, "Berhasil memperbarui kategori", data)
}
