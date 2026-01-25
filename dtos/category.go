package dtos

import "time"

// CategoryCreateRequest represents the request body for creating a category
type CategoryCreateRequest struct {
	Name        string `json:"name" example:"Electronics"`
	Description string `json:"description" example:"Electronic items"`
	IsActive    bool   `json:"is_active" example:"true"`
}

// CategoryUpdateRequest represents the request body for updating a category
type CategoryUpdateRequest struct {
	Name        string `json:"name" example:"Electronics"`
	Description string `json:"description" example:"Electronic items"`
	IsActive    bool   `json:"is_active" example:"true"`
}

// CategoryResponse represents the response body for category
type CategoryResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Electronics"`
	Description string    `json:"description" example:"Electronic items"`
	IsActive    bool      `json:"is_active" example:"true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateCategorySuccessResponse represents the success response for create category
type CreateCategorySuccessResponse struct {
	Message string           `json:"message" example:"Category created successfully"`
	Data    CategoryResponse `json:"data"`
}

// GetCategoriesSuccessResponse represents the success response for get all categories
type GetCategoriesSuccessResponse struct {
	Message string             `json:"message" example:"Categories retrieved successfully"`
	Data    []CategoryResponse `json:"data"`
	Total   int                `json:"total" example:"1"`
}

// GetCategorySuccessResponse represents the success response for get category by ID
type GetCategorySuccessResponse struct {
	Message string           `json:"message" example:"Category retrieved successfully"`
	Data    CategoryResponse `json:"data"`
}

// UpdateCategorySuccessResponse represents the success response for update category
type UpdateCategorySuccessResponse struct {
	Message string           `json:"message" example:"Category updated successfully"`
	Data    CategoryResponse `json:"data"`
}

// DeleteCategorySuccessResponse represents the success response for delete category
type DeleteCategorySuccessResponse struct {
	Message string `json:"message" example:"Category deleted successfully"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}
