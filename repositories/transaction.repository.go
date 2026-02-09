package repositories

import (
	"errors"
	"time"

	"kasir-api/models"

	"gorm.io/gorm"
)

// ITransactionRepository defines the interface for transaction data operations
type ITransactionRepository interface {
	CreateTransaction() (*models.Transactions, error)
	GetTransactionByID(id int) (*models.Transactions, error)
	AddItemToTransaction(transactionID, productID, quantity int, productName string, subtotal int) error
	UpdateTransactionTotal(transactionID int, totalAmount int) error
	GetTransactionsByDateRange(start, end time.Time) ([]models.Transactions, error)
	GetReportByDateRange(start, end time.Time) ([]map[string]interface{}, error)
}

// TransactionRepository handles database operations for transactions
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new transaction repository instance
func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

// CreateTransaction creates a new transaction
func (r *TransactionRepository) CreateTransaction() (*models.Transactions, error) {
	transaction := &models.Transactions{
		TotalAmount: 0,
	}

	result := r.db.Create(transaction)
	return transaction, result.Error
}

// GetTransactionByID retrieves a transaction by ID with its details
func (r *TransactionRepository) GetTransactionByID(id int) (*models.Transactions, error) {
	var transaction models.Transactions

	result := r.db.
		Preload("Details").
		First(&transaction, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, result.Error
	}

	return &transaction, nil
}

// AddItemToTransaction adds an item (product) to a transaction
func (r *TransactionRepository) AddItemToTransaction(transactionID, productID, quantity int, productName string, subtotal int) error {
	detail := models.TransactionDetails{
		TransactionID: transactionID,
		ProductID:     productID,
		ProductName:   productName,
		Quantity:      quantity,
		Subtotal:      subtotal,
	}

	result := r.db.Create(&detail)
	return result.Error
}

// UpdateTransactionTotal updates the total amount of a transaction
func (r *TransactionRepository) UpdateTransactionTotal(transactionID int, totalAmount int) error {
	result := r.db.Model(&models.Transactions{}).
		Where("id = ?", transactionID).
		Update("total_amount", totalAmount)

	return result.Error
}

// GetTransactionsByDateRange retrieves transactions within provided date range (inclusive)
func (r *TransactionRepository) GetTransactionsByDateRange(start, end time.Time) ([]models.Transactions, error) {
	var transactions []models.Transactions

	result := r.db.
		Preload("Details").
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("created_at desc").
		Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

// GetReportByDateRange retrieves best-selling products within provided date range
func (r *TransactionRepository) GetReportByDateRange(start, end time.Time) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// Query to get best-selling products with their total quantities sold
	queryResult := r.db.
		Table("transaction_details").
		Select("product_name as name, SUM(quantity) as qty_sold").
		Where("created_at BETWEEN ? AND ?", start, end).
		Group("product_name").
		Order("qty_sold DESC").
		Find(&result)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return result, nil
}
