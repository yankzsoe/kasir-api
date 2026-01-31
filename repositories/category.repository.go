package repositories

import (
	"errors"

	"kasir-api/models"

	"gorm.io/gorm"
)

// ICategoryRepository defines the interface for category data operations
type ICategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id string) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Update(id string, updateData *models.Category) (*models.Category, error)
	Delete(id string) error
}

// CategoryRepository handles database operations for categories
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository instance
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// Create creates a new category in the database
func (r *CategoryRepository) Create(category *models.Category) error {
	result := r.db.Create(category)
	return result.Error
}

// FindByID retrieves a single category by ID
func (r *CategoryRepository) FindByID(id string) (*models.Category, error) {
	var category models.Category
	result := r.db.Find(&category, id)
	return &category, result.Error
}

// FindAll retrieves all categories from the database
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}

// Update updates an existing category
func (r *CategoryRepository) Update(id string, updateData *models.Category) (*models.Category, error) {
	var category models.Category

	// Find category by ID
	result := r.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, result.Error
	}

	// Update category
	result = r.db.Model(&category).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

// Delete deletes a category by ID (soft delete)
func (r *CategoryRepository) Delete(id string) error {
	var category models.Category

	// Find category by ID
	result := r.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return result.Error
	}

	// Delete category (soft delete via DeletedAt)
	result = r.db.Delete(&category)
	return result.Error
}
