package repositories

import "gorm.io/gorm"

// Repository holds all repository instances
type Repository struct {
	Category ICategoryRepository
}

// NewRepository initializes all repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
	}
}
