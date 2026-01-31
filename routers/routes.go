package routers

import (
	"kasir-api/controllers"
	_ "kasir-api/docs"
	"kasir-api/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(router *gin.Engine) {
	router.Use(middlewares.ErrorHandler())
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

	// Product routes
	setupProductRoutes(router)
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

// setupProductRoutes configures product-related routes
func setupProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/api/v1/products")
	{
		// Create product
		// @Summary Create a new product
		// @Description Create a new product with the provided details
		// @Tags Products
		// @Accept json
		// @Produce json
		// @Param product body models.Produk true "Product data"
		// @Success 201 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /products [post]
		productGroup.POST("", controllers.CreateProduct)

		// Get all products
		// @Summary Get all products
		// @Description Retrieve a list of all products
		// @Tags Products
		// @Produce json
		// @Success 200 {object} map[string]interface{}
		// @Failure 500 {object} map[string]string
		// @Router /products [get]
		productGroup.GET("", controllers.GetAllProducts)

		// Get product by ID
		// @Summary Get a product by ID
		// @Description Retrieve a specific product by its ID
		// @Tags Products
		// @Produce json
		// @Param id path string true "Product ID"
		// @Success 200 {object} map[string]interface{}
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /products/{id} [get]
		productGroup.GET("/:id", controllers.GetProductByID)

		// Update product
		// @Summary Update a product
		// @Description Update an existing product by ID
		// @Tags Products
		// @Accept json
		// @Produce json
		// @Param id path string true "Product ID"
		// @Param product body models.Produk true "Product data to update"
		// @Success 200 {object} map[string]interface{}
		// @Failure 400 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /products/{id} [put]
		productGroup.PUT("/:id", controllers.UpdateProduct)

		// Delete product
		// @Summary Delete a product
		// @Description Delete a product by ID (soft delete)
		// @Tags Products
		// @Produce json
		// @Param id path string true "Product ID"
		// @Success 200 {object} map[string]string
		// @Failure 404 {object} map[string]string
		// @Failure 500 {object} map[string]string
		// @Router /products/{id} [delete]
		productGroup.DELETE("/:id", controllers.DeleteProduct)
	}
}
