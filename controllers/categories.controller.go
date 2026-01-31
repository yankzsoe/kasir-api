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

var categoryService *services.CategoryService

// SetCategoryService initializes the category service with repository
func SetCategoryService(service *services.CategoryService) {
	categoryService = service
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body dtos.CategoryCreateRequest true "Category data"
// @Success 201 {object} dtos.CreateCategorySuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var req dtos.CategoryCreateRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Convert DTO to model
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	// Call service to create category
	if err := categoryService.CreateCategory(&category); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			common.ThrowException(400, "bad request (the name can't duplicate)")
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert model to response DTO
	response := dtos.CreateCategorySuccessResponse{
		Message: "Category created successfully",
		Data: dtos.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags Categories
// @Produce json
// @Success 200 {object} dtos.GetCategoriesSuccessResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	// Call service to get all categories
	categories, err := categoryService.GetAllCategories()
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert models to response DTOs
	categoryResponses := make([]dtos.CategoryResponse, len(categories))
	for i, cat := range categories {
		categoryResponses[i] = dtos.CategoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
			IsActive:    cat.IsActive,
			CreatedAt:   cat.CreatedAt,
			UpdatedAt:   cat.UpdatedAt,
		}
	}

	response := common.GetListSuccessResponse(categoryResponses)

	c.JSON(http.StatusOK, response)
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Retrieve a specific category by its ID
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} dtos.GetCategorySuccessResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	uriReq := dtos.CategoryUriRequest{}
	if err := c.ShouldBindUri(&uriReq); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Call service to get category by ID
	category, err := categoryService.GetCategoryByID(uriReq.ID)
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if category.ID == 0 {
		response := common.NotFoundResponse()
		c.JSON(http.StatusOK, response)
		return
	}
	response := common.SuccessResponseWithData(category)

	c.JSON(http.StatusOK, response)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body dtos.CategoryUpdateRequest true "Category data to update"
// @Success 200 {object} dtos.UpdateCategorySuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	var req dtos.CategoryUpdateRequest
	uriReq := dtos.CategoryUriRequest{}

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
	updateData := models.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	// Call service to update category
	category, err := categoryService.UpdateCategory(uriReq.ID, &updateData)
	if err != nil {
		if err.Error() == "category not found" {
			common.ThrowException(http.StatusNotFound, err.Error())
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert model to response DTO
	response := dtos.UpdateCategorySuccessResponse{
		Message: "Category updated successfully",
		Data: dtos.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID (soft delete)
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} dtos.DeleteCategorySuccessResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	uriReq := dtos.CategoryUriRequest{}
	if err := c.ShouldBindUri(&uriReq); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Call service to delete category
	if err := categoryService.DeleteCategory(uriReq.ID); err != nil {
		if err.Error() == "category not found" {
			common.ThrowException(http.StatusNotFound, err.Error())
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	response := dtos.DeleteCategorySuccessResponse{
		Message: "Category deleted successfully",
	}

	c.JSON(http.StatusOK, response)
}
