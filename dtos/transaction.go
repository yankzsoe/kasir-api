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

// TransactionDetailResponse represents a detail item in transaction list
type TransactionDetailResponse struct {
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Subtotal    int       `json:"subtotal"`
	CreatedAt   time.Time `json:"created_at"`
}

// TransactionResponse represents a transaction in list response
type TransactionResponse struct {
	TransactionID int                         `json:"transaction_id"`
	TotalAmount   int                         `json:"total_amount"`
	Details       []TransactionDetailResponse `json:"details"`
	CreatedAt     time.Time                   `json:"created_at"`
}

// BestSellingProduct represents a best-selling product in the report
type BestSellingProduct struct {
	Name    string `json:"name" example:"Indomie Goreng 1"`
	QtySold int    `json:"qty_sold" example:"12"`
}

// TransactionReportResponse represents the selling report response
type TransactionReportResponse struct {
	ReportDate          string               `json:"report_date" example:"2026-01-15"`
	TotalRevenue        int                  `json:"total_revenue" example:"45000"`
	TotalTransactions   int                  `json:"total_transactions" example:"5"`
	BestSellingProducts []BestSellingProduct `json:"best-selling-product"`
}
