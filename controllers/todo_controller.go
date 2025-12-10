package controllers

import (
	"fmt"
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	var req dto.CreateDtoTodo

	// debug: print engine saat request masuk
	fmt.Printf("Request incoming — binding.Engine: %T\n", binding.Validator.Engine())

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// debug: tunjukkan tipe error
		fmt.Println("ShouldBindJSON error:", err)

		// cek apakah error adalah ValidationErrors
		if ve, ok := err.(validator.ValidationErrors); ok {
			fmt.Println("ValidationErrors detected:")
			for _, e := range ve {
				fmt.Printf("- Field: %s, Tag: %s, Param: %s\n", e.Field(), e.Tag(), e.Param())
			}
		}
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
