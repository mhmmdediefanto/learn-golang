package main

import (
	"go-bakcend-todo-list/config"
	"go-bakcend-todo-list/controllers"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
	"go-bakcend-todo-list/routes"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	utils.InitValidator()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	// DB
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	config.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.Category{},
		&models.Article{},
	)

	log.Println("Databases Connected")

	r := gin.Default()

	// ===== REPOSITORY =====
	userRepo := repositories.NewUserRepository(config.DB)
	todoRepo := repositories.NewTodoRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	articleRepo := repositories.NewArticleRepository(config.DB)

	// ===== SERVICE =====
	userService := services.NewUserService(userRepo)
	todoService := services.NewTodoService(todoRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	authService := services.NewAuthService(userRepo)
	articleService := services.NewArticleService(articleRepo)

	// ===== CONTROLLER =====
	userController := controllers.NewUserController(userService)
	todoController := controllers.NewTodoController(todoService)
	categoryController := controllers.NewCategoryController(categoryService)
	authController := controllers.NewAuthController(authService)
	articleController := controllers.NewArticleController(articleService)

	// ===== ROUTES =====
	routes.SetupRoutes(
		r,
		userController,
		todoController,
		categoryController,
		authController,
		authService,
		articleController,
	)

	r.Run(":8080")
}
