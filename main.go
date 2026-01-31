package main

import (
	"log"
	"os"

	"kasir-api/configs"
	"kasir-api/controllers"
	"kasir-api/docs"
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/routers"
	"kasir-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Kasir API
// @version 1.0
// @description A complete Point of Sale (POS) API built with Go, Gin, and GORM
// @host localhost:8080
// @basePath /api/v1
// @schemes http
// @contact.name API Support
// @contact.url http://www.kasir-api.com/support
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Load configurations
	cfg := configs.LoadConfig()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	if err := configs.InitDB(cfg.Database.Path); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate models
	db := configs.GetDB()
	if err := db.AutoMigrate(&models.Category{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Initialize repositories
	repo := repositories.NewRepository(db)

	// Initialize services
	categoryService := services.NewCategoryService(repo.Category)

	// Set up controllers with services
	controllers.SetCategoryService(categoryService)

	// Create Gin router
	router := gin.Default()

	// Setup routes
	routers.SetupRoutes(router)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Kasir API"
	docs.SwaggerInfo.Description = "A complete Point of Sale (POS) API built with Go, Gin, and GORM"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	// Start server
	log.Printf("Starting Kasir API v%s on port %s...", cfg.API.Version, cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
