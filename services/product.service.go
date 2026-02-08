package services

import (
	"errors"

	"kasir-api/models"
	"kasir-api/repositories"
)

// ProductService handles business logic for products
type ProductService struct {
	repo repositories.IProductRepository
}

// NewProductService creates a new product service instance
func NewProductService(repo repositories.IProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// CreateProduct creates a new product in the database
func (s *ProductService) CreateProduct(product *models.Produk) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}

	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.repo.Create(product)
}

// GetAllProducts retrieves all products from the database
func (s *ProductService) GetAllProducts() ([]models.Produk, error) {
	return s.repo.FindAll()
}

// SearchProducts searches for products by name or category name
func (s *ProductService) SearchProducts(query string) ([]models.Produk, error) {
	if query == "" {
		return s.repo.FindAll()
	}
	return s.repo.Search(query)
}

// GetProductByID retrieves a single product by ID
func (s *ProductService) GetProductByID(id string) (*models.Produk, error) {
	return s.repo.FindByID(id)
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id string, updateData *models.Produk) (*models.Produk, error) {
	// Validate required fields
	if updateData.Name == "" {
		// Get current product to check if it has a name
		current, err := s.repo.FindByID(id)
		if err != nil {
			return nil, err
		}
		if current.Name == "" {
			return nil, errors.New("product name is required")
		}
	}

	if updateData.Price < 0 {
		return nil, errors.New("product price cannot be negative")
	}

	if updateData.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	return s.repo.Update(id, updateData)
}

// DeleteProduct deletes a product by ID (soft delete)
func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}
