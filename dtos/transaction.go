package dtos

import "time"

// CheckoutItem represents a product item for checkout
type CheckoutItem struct {
	ProductID int `json:"product_id" validate:"required" example:"1"`
	Quantity  int `json:"quantity" validate:"required,min=1" example:"2"`
}

// CheckoutRequest represents the request body for completing a transaction
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items" validate:"required,min=1"`
}

// CheckoutItemResponse represents a product item in the checkout response
type CheckoutItemResponse struct {
	ProductID   int    `json:"product_id" example:"1"`
	ProductName string `json:"product_name" example:"Laptop"`
	Quantity    int    `json:"quantity" example:"2"`
	Price       int    `json:"price" example:"10000000"`
	Subtotal    int    `json:"subtotal" example:"20000000"`
}

// CheckoutResponse represents the checkout response with transaction details
type CheckoutResponse struct {
	TransactionID int                    `json:"transaction_id" example:"1"`
	Items         []CheckoutItemResponse `json:"items"`
	TotalAmount   int                    `json:"total_amount" example:"20000000"`
	CreatedAt     time.Time              `json:"created_at"`
}

// CheckoutSuccessResponse represents the response after completing checkout
type CheckoutSuccessResponse struct {
	Message string           `json:"message" example:"Checkout completed successfully"`
	Data    CheckoutResponse `json:"data"`
}
