package routes

import (
	"go-bakcend-todo-list/controllers"
	"go-bakcend-todo-list/middleware"
	"go-bakcend-todo-list/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	todoHandler := controllers.TodoController{}
	userHandler := controllers.UserController{}
	authHandler := controllers.AuthController{}
	userRepo := repositories.NewUserRepository()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running 🚀",
		})
	})

	// 3. API Grouping & Versioning
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/me", middleware.AuthMiddleware(), authHandler.Me)
			auth.POST("/signin", authHandler.SignIn)
			auth.POST("/refresh", middleware.RefreshTokenMiddleware(userRepo), authHandler.RefreshToken)
			// auth.POST("/register", userHandler.Create) // Register user baru biasanya public
			auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
		}

		// User Routes
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetAll)
			users.POST("/", userHandler.Create)
		}

		// Todo Routes
		todos := api.Group("/todos", middleware.AuthMiddleware())
		{
			todos.GET("/", todoHandler.GetAll)
			todos.POST("/", todoHandler.Create)
			// todos.GET("/:id", todoHandler.GetByID)
			todos.PATCH("/:id", todoHandler.Update)
			todos.DELETE("/:id", todoHandler.Delete)
		}

	}
}
