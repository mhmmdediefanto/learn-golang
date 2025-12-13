package controllers

import (
	"fmt"
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service services.TodoService
}

func (c *TodoController) GetAll(ctx *gin.Context) {
	todos, err := c.service.GetAll()
	if err != nil {
		utils.Error(ctx, 500, "Gagal Mengambil Todo List", err)
		return
	}

	message := "Berhasil mengambil todo"
	if len(todos) == 0 {
		message = "Data todo kosong"
	}

	utils.Success(ctx, message, todos)
}

func (c *TodoController) Create(ctx *gin.Context) {
	var req dto.CreateDtoTodo

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	todo := models.Todo{
		Title:  req.Title,
		Status: req.Status,
	}

	if err := c.service.Create(&todo); err != nil {
		utils.Error(ctx, 500, "Gagal menyimpan todo", err)
		return
	}

	utils.Created(ctx, "Todo berhasil dibuat", todo)
}

func (c *TodoController) Delete(ctx *gin.Context) {
	idParams := ctx.Param("id")

	id64, err := strconv.ParseUint(idParams, 10, 32)
	if err != nil {
		utils.Error(ctx, 400, "ID tidak valid", err)
		return
	}

	id := uint(id64)
	if err := c.service.Delete(id); err != nil {
		utils.Error(ctx, 404, "Todo tidak ditemukan", err)
		return
	}

	utils.Success(ctx, "Todo berhasil dihapus", nil)

}

func (c *TodoController) Update(ctx *gin.Context) {

	id, err := utils.ParamUint(ctx, "id")
	if err != nil {
		utils.Error(ctx, 400, "ID tidak valid", err)
		return
	}

	fmt.Println("ID yang diupdate:", id)

	// Bind JSON request ke DTO
	var req dto.CreateDtoTodo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// Convert DTO ke model
	todo := &models.Todo{
		Title:  req.Title,
		Status: req.Status,
	}
	// Panggil service update
	updatedTodo, err := c.service.Update(id, todo)
	if err != nil {
		utils.Error(ctx, 500, "Gagal mengupdate todo", err)
		return
	}
	utils.Success(ctx, "Todo berhasil diupdate", updatedTodo)
}
