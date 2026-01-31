package dtos

import "time"

// ProductCreateRequest represents the request body for creating a product
type ProductCreateRequest struct {
	Name       string `json:"name" validate:"required" example:"Laptop"`
	Price      int    `json:"price" validate:"required" example:"10000000"`
	Stock      int    `json:"stock" validate:"required" example:"5"`
	IsActive   bool   `json:"is_active" example:"true"`
	CategoryId *int   `json:"category_id" example:"1"`
}

// ProductUpdateRequest represents the request body for updating a product
type ProductUpdateRequest struct {
	Name       string `json:"name" example:"Laptop"`
	Price      int    `json:"price" example:"10000000"`
	Stock      int    `json:"stock" example:"5"`
	IsActive   bool   `json:"is_active" example:"true"`
	CategoryId *int   `json:"category_id" example:"1"`
}

// ProductUriRequest represents the URI parameter for product ID
type ProductUriRequest struct {
	ID string `uri:"id" validate:"required" example:"1"`
}

// ProductResponse represents the response body for product
type ProductResponse struct {
	ID           int       `json:"id" example:"1"`
	Name         string    `json:"name" example:"Laptop"`
	CategoryName string    `json:"category_name" example:"Electronic"`
	Price        int       `json:"price" example:"10000000"`
	Stock        int       `json:"stock" example:"5"`
	IsActive     bool      `json:"is_active" example:"true"`
	CategoryId   *int      `json:"category_id" example:"1"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateProductResponse struct {
	ID         int       `json:"id" example:"1"`
	Name       string    `json:"name" example:"Laptop"`
	Price      int       `json:"price" example:"10000000"`
	Stock      int       `json:"stock" example:"5"`
	IsActive   bool      `json:"is_active" example:"true"`
	CategoryId *int      `json:"category_id" example:"1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateProductSuccessResponse represents the success response for create product
type CreateProductSuccessResponse struct {
	Message string          `json:"message" example:"Product created successfully"`
	Data    ProductResponse `json:"data"`
}

// GetProductsSuccessResponse represents the success response for get all products
type GetProductsSuccessResponse struct {
	Message string            `json:"message" example:"Products retrieved successfully"`
	Data    []ProductResponse `json:"data"`
	Total   int               `json:"total" example:"1"`
}

// GetProductSuccessResponse represents the success response for get product by ID
type GetProductSuccessResponse struct {
	Message string          `json:"message" example:"Product retrieved successfully"`
	Data    ProductResponse `json:"data"`
}

// UpdateProductSuccessResponse represents the success response for update product
type UpdateProductSuccessResponse struct {
	Message string                `json:"message" example:"Product updated successfully"`
	Data    UpdateProductResponse `json:"data"`
}

// DeleteProductSuccessResponse represents the success response for delete product
type DeleteProductSuccessResponse struct {
	Message string `json:"message" example:"Product deleted successfully"`
}
