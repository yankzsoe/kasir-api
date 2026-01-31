package repositories

import (
	"errors"

	"kasir-api/models"

	"gorm.io/gorm"
)

// IProductRepository defines the interface for product data operations
type IProductRepository interface {
	Create(product *models.Produk) error
	FindByID(id string) (*models.Produk, error)
	FindAll() ([]models.Produk, error)
	Update(id string, updateData *models.Produk) (*models.Produk, error)
	Delete(id string) error
}

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Create creates a new product in the database
func (r *ProductRepository) Create(product *models.Produk) error {
	result := r.db.Create(product)
	return result.Error
}

// FindByID retrieves a single product by ID
func (r *ProductRepository) FindByID(id string) (*models.Produk, error) {
	var product models.Produk
	result := r.db.Find(&product, id)
	return &product, result.Error
}

// FindAll retrieves all products from the database
func (r *ProductRepository) FindAll() ([]models.Produk, error) {
	var products []models.Produk
	result := r.db.Find(&products)
	return products, result.Error
}

// Update updates an existing product
func (r *ProductRepository) Update(id string, updateData *models.Produk) (*models.Produk, error) {
	var product models.Produk

	// Find product by ID
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}

	// Update product
	result = r.db.Model(&product).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// Delete deletes a product by ID (soft delete)
func (r *ProductRepository) Delete(id string) error {
	var product models.Produk

	// Find product by ID
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return result.Error
	}

	// Delete product (soft delete via DeletedAt)
	result = r.db.Delete(&product)
	return result.Error
}
