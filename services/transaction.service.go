package services

import (
	"errors"
	"strconv"
	"time"

	"kasir-api/dtos"
	"kasir-api/models"
	"kasir-api/repositories"
)

// TransactionService handles business logic for transactions
type TransactionService struct {
	transactionRepo repositories.ITransactionRepository
	productRepo     repositories.IProductRepository
}

// NewTransactionService creates a new transaction service instance
func NewTransactionService(transactionRepo repositories.ITransactionRepository, productRepo repositories.IProductRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
	}
}

// ProcessCheckout processes a checkout transaction with items, creates transaction and reduces stock
func (s *TransactionService) ProcessCheckout(items []dtos.CheckoutItem) (*models.Transactions, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided for checkout")
	}

	// Create a new transaction
	transaction, err := s.transactionRepo.CreateTransaction()
	if err != nil {
		return nil, err
	}

	totalAmount := 0

	// Process each item
	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than 0")
		}

		// Get product details to validate and get price
		product, err := s.productRepo.FindByID(strconv.Itoa(item.ProductID))
		if err != nil {
			return nil, err
		}

		if product.ID == 0 {
			return nil, errors.New("product not found")
		}

		// Check if product has sufficient stock
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		// Calculate subtotal
		subtotal := product.Price * item.Quantity

		// Add item to transaction details
		err = s.transactionRepo.AddItemToTransaction(transaction.ID, product.ID, item.Quantity, product.Name, subtotal)
		if err != nil {
			return nil, err
		}

		// Update product stock
		updatedProduct := &models.Produk{
			Stock: product.Stock - item.Quantity,
		}

		_, err = s.productRepo.Update(strconv.Itoa(product.ID), updatedProduct)
		if err != nil {
			return nil, err
		}

		// Update total amount
		totalAmount += subtotal
	}

	// Update transaction total
	err = s.transactionRepo.UpdateTransactionTotal(transaction.ID, totalAmount)
	if err != nil {
		return nil, err
	}

	// Return the completed transaction with all its details
	return s.transactionRepo.GetTransactionByID(transaction.ID)
}

// SearchTransactionsByDateRange returns transactions within start and end time
func (s *TransactionService) SearchTransactionsByDateRange(start, end time.Time) ([]models.Transactions, error) {
	return s.transactionRepo.GetTransactionsByDateRange(start, end)
}

// GenerateReport generates a selling report for a date range
func (s *TransactionService) GenerateReport(start, end time.Time) (*dtos.TransactionReportResponse, error) {
	// Get transactions for revenue and count calculation
	transactions, err := s.transactionRepo.GetTransactionsByDateRange(start, end)
	if err != nil {
		return nil, err
	}

	// Calculate total revenue and count
	totalRevenue := 0
	for _, t := range transactions {
		totalRevenue += t.TotalAmount
	}

	// Get best-selling products
	bestSellingData, err := s.transactionRepo.GetReportByDateRange(start, end)
	if err != nil {
		return nil, err
	}

	// Convert to BestSellingProduct DTO, limit to top 3
	bestSellingProducts := make([]dtos.BestSellingProduct, 0)
	for i, item := range bestSellingData {
		if i >= 3 {
			break
		}

		// Handle name safely
		var productName string
		if name, ok := item["name"].(string); ok {
			productName = name
		}

		// Handle qty_sold which might be int64 or float64 depending on database driver
		var qtySold int
		switch v := item["qty_sold"].(type) {
		case int64:
			qtySold = int(v)
		case float64:
			qtySold = int(v)
		case int:
			qtySold = v
		default:
			qtySold = 0
		}

		bestSellingProducts = append(bestSellingProducts, dtos.BestSellingProduct{
			Name:    productName,
			QtySold: qtySold,
		})
	}

	// Format report date as date range
	reportDateStr := start.Format("2006-01-02") + " to " + end.Format("2006-01-02")

	report := &dtos.TransactionReportResponse{
		ReportDate:          reportDateStr,
		TotalRevenue:        totalRevenue,
		TotalTransactions:   len(transactions),
		BestSellingProducts: bestSellingProducts,
	}

	return report, nil
}
