package services

import (
	"errors"
	"strconv"

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
