package routes

import (
	"go-bakcend-todo-list/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	todo := controllers.TodoController{}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running 🚀",
		})
	})

	r.GET("/todos", todo.GetAllTodos)
	r.POST("/todos", todo.CreateTodo)
	r.DELETE("/todos/:id", todo.DeleteTodo)

}
