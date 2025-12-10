package controllers

import (
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service services.TodoService
}

func (c *TodoController) GetAllTodos(ctx *gin.Context) {
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

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo

	if err := ctx.ShouldBind(&todo); err != nil {
		utils.Error(ctx, 400, "Invalid request payload", err)
		return
	}

	if err := c.service.Create(&todo); err != nil {
		utils.Error(ctx, 500, "Gagal menyimpan todo", err)
		return
	}

	utils.Created(ctx, "Todo berhasil dibuat", todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
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
