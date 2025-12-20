package routes

import (
	"go-bakcend-todo-list/controllers"
	"go-bakcend-todo-list/middleware"
	"go-bakcend-todo-list/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userController *controllers.UserController,
	todoController *controllers.TodoController,
	categoryController *controllers.CategoryController,
	authController *controllers.AuthController,
	authService services.AuthService,
	articleController *controllers.ArticleController,
) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running 🚀",
		})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/me", middleware.AuthMiddleware(), authController.Me)
			auth.POST("/signin", authController.SignIn)
			auth.POST(
				"/refresh",
				middleware.RefreshTokenMiddleware(authService),
				authController.RefreshToken,
			)
			auth.POST("/logout", middleware.AuthMiddleware(), authController.Logout)
		}

		users := api.Group("/users")
		{
			users.GET("/", userController.GetAll)
			users.POST("/", userController.Create)
		}

		todos := api.Group("/todos", middleware.AuthMiddleware())
		{
			todos.GET("/", todoController.GetAll)
			todos.POST("/", todoController.Create)
			todos.PATCH("/:id", todoController.Update)
			todos.DELETE("/:id", todoController.Delete)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", categoryController.GetAll)
			categories.POST("/", categoryController.Create)
			categories.PUT("/:id", categoryController.Update)
			categories.DELETE("/:id", categoryController.Delete)
		}

		article := api.Group("/articles", middleware.AuthMiddleware())
		{
			article.GET("/", articleController.GetAll)
			article.POST("/", articleController.Create)
			article.DELETE("/:id", articleController.Delete)
			article.PUT("/:id", articleController.Update)
		}
	}
}
