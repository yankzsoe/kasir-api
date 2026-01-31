package repositories

import "gorm.io/gorm"

// Repository holds all repository instances
type Repository struct {
	Category ICategoryRepository
	Product  IProductRepository
}

// NewRepository initializes all repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		Product:  NewProductRepository(db),
	}
}
