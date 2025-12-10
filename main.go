package main

import (
	"go-bakcend-todo-list/config"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{})

	// jika success kasih info log
	log.Println("Databases Connected")

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
