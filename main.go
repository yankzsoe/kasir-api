package main

import (
	"log"

	"kasir-api/configs"
	"kasir-api/models"
	"kasir-api/routers"

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

	// Create Gin router
	router := gin.Default()

	// Setup routes
	routers.SetupRoutes(router)

	// Start server
	log.Printf("Starting Kasir API v%s on port %s...", cfg.API.Version, cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
