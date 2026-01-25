package routers

import (
	"kasir-api/controllers"
	_ "kasir-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(router *gin.Engine) {
	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check route
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
			"status":  "healthy",
		})
	})

	// Category routes
	setupCategoryRoutes(router)
}

// setupCategoryRoutes configures category-related routes
func setupCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/api/v1/categories")
	{
		// Create category
		// @Summary Create a new category
		// @Description Create a new category with the provided details
		// @Tags Categories
		// @Accept json
		// @Produce json
		// @Param category body models.Category true "Category data"
		// @Success 201 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /categories [post]
		categoryGroup.POST("", controllers.CreateCategory)

		// Get all categories
		// @Summary Get all categories
		// @Description Retrieve a list of all categories
		// @Tags Categories
		// @Produce json
		// @Success 200 {object} map[string]interface{}
		// @Failure 500 {object} map[string]string
		// @Router /categories [get]
		categoryGroup.GET("", controllers.GetAllCategories)

		// Get category by ID
		// @Summary Get a category by ID
		// @Description Retrieve a specific category by its ID
		// @Tags Categories
		// @Produce json
		// @Param id path string true "Category ID"
		// @Success 200 {object} map[string]interface{}
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /categories/{id} [get]
		categoryGroup.GET("/:id", controllers.GetCategoryByID)

		// Update category
		// @Summary Update a category
		// @Description Update an existing category by ID
		// @Tags Categories
		// @Accept json
		// @Produce json
		// @Param id path string true "Category ID"
		// @Param category body models.Category true "Category data to update"
		// @Success 200 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /categories/{id} [put]
		categoryGroup.PUT("/:id", controllers.UpdateCategory)

		// Delete category
		// @Summary Delete a category
		// @Description Delete a category by ID (soft delete)
		// @Tags Categories
		// @Produce json
		// @Param id path string true "Category ID"
		// @Success 200 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /categories/{id} [delete]
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}
}
