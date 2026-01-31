package controllers

import (
	"net/http"
	"strings"

	"kasir-api/common"
	"kasir-api/dtos"
	"kasir-api/models"
	"kasir-api/services"

	"github.com/gin-gonic/gin"
)

var productService *services.ProductService

// SetProductService initializes the product service with repository
func SetProductService(service *services.ProductService) {
	productService = service
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dtos.ProductCreateRequest true "Product data"
// @Success 201 {object} dtos.CreateProductSuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var req dtos.ProductCreateRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Convert DTO to model
	product := models.Produk{
		Nama:       req.Nama,
		Harga:      req.Harga,
		Stok:       req.Stok,
		IsActive:   req.IsActive,
		CategoryId: req.CategoryId,
	}

	// Call service to create product
	if err := productService.CreateProduct(&product); err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert model to response DTO
	response := dtos.CreateProductSuccessResponse{
		Message: "Product created successfully",
		Data: dtos.ProductResponse{
			ID:         product.ID,
			Nama:       product.Nama,
			Harga:      product.Harga,
			Stok:       product.Stok,
			IsActive:   product.IsActive,
			CategoryId: product.CategoryId,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags Products
// @Produce json
// @Success 200 {object} dtos.GetProductsSuccessResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	// Call service to get all products
	products, err := productService.GetAllProducts()
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert models to response DTOs
	productResponses := make([]dtos.ProductResponse, len(products))
	for i, prod := range products {
		productResponses[i] = dtos.ProductResponse{
			ID:         prod.ID,
			Nama:       prod.Nama,
			Harga:      prod.Harga,
			Stok:       prod.Stok,
			IsActive:   prod.IsActive,
			CategoryId: prod.CategoryId,
			CreatedAt:  prod.CreatedAt,
			UpdatedAt:  prod.UpdatedAt,
		}
	}

	response := common.GetListSuccessResponse(productResponses)

	c.JSON(http.StatusOK, response)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieve a specific product by its ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dtos.GetProductSuccessResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	uriReq := dtos.ProductUriRequest{}
	if err := c.ShouldBindUri(&uriReq); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Call service to get product by ID
	product, err := productService.GetProductByID(uriReq.ID)
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if product.ID == 0 {
		response := common.NotFoundResponse()
		c.JSON(http.StatusOK, response)
		return
	}

	response := dtos.GetProductSuccessResponse{
		Message: "Product retrieved successfully",
		Data: dtos.ProductResponse{
			ID:         product.ID,
			Nama:       product.Nama,
			Harga:      product.Harga,
			Stok:       product.Stok,
			IsActive:   product.IsActive,
			CategoryId: product.CategoryId,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body dtos.ProductUpdateRequest true "Product data to update"
// @Success 200 {object} dtos.UpdateProductSuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var req dtos.ProductUpdateRequest
	uriReq := dtos.ProductUriRequest{}

	// Bind and validate URI parameters
	if err := c.ShouldBindUri(&uriReq); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Convert DTO to model
	updateData := models.Produk{
		Nama:       req.Nama,
		Harga:      req.Harga,
		Stok:       req.Stok,
		IsActive:   req.IsActive,
		CategoryId: req.CategoryId,
	}

	// Call service to update product
	product, err := productService.UpdateProduct(uriReq.ID, &updateData)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			common.ThrowException(http.StatusNotFound, err.Error())
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert model to response DTO
	response := dtos.UpdateProductSuccessResponse{
		Message: "Product updated successfully",
		Data: dtos.ProductResponse{
			ID:         product.ID,
			Nama:       product.Nama,
			Harga:      product.Harga,
			Stok:       product.Stok,
			IsActive:   product.IsActive,
			CategoryId: product.CategoryId,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID (soft delete)
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dtos.DeleteProductSuccessResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	uriReq := dtos.ProductUriRequest{}
	if err := c.ShouldBindUri(&uriReq); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Call service to delete product
	if err := productService.DeleteProduct(uriReq.ID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			common.ThrowException(http.StatusNotFound, err.Error())
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	response := dtos.DeleteProductSuccessResponse{
		Message: "Product deleted successfully",
	}

	c.JSON(http.StatusOK, response)
}
