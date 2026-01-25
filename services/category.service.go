package services

import (
	"errors"

	"kasir-api/configs"
	"kasir-api/models"

	"gorm.io/gorm"
)

// CategoryService handles business logic for categories
type CategoryService struct {
	db *gorm.DB
}

// NewCategoryService creates a new category service instance
func NewCategoryService() *CategoryService {
	return &CategoryService{
		db: configs.GetDB(),
	}
}

// CreateCategory creates a new category in the database
func (s *CategoryService) CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	db := configs.GetDB()
	result := db.Create(category)
	return result.Error
}

// GetAllCategories retrieves all categories from the database
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	db := configs.GetDB()
	result := db.Find(&categories)
	return categories, result.Error
}

// GetCategoryByID retrieves a single category by ID
func (s *CategoryService) GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	db := configs.GetDB()
	result := db.First(&category, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("category not found")
		}
		return nil, result.Error
	}

	return &category, nil
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(id string, updateData *models.Category) (*models.Category, error) {
	var category models.Category
	db := configs.GetDB()

	// Find category by ID
	result := db.First(&category, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("category not found")
		}
		return nil, result.Error
	}

	// Validate required fields
	if updateData.Name == "" && category.Name == "" {
		return nil, errors.New("category name is required")
	}

	// Update category
	result = db.Model(&category).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

// DeleteCategory deletes a category by ID (soft delete)
func (s *CategoryService) DeleteCategory(id string) error {
	var category models.Category
	db := configs.GetDB()

	// Find category by ID
	result := db.First(&category, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("category not found")
		}
		return result.Error
	}

	// Delete category (soft delete via DeletedAt)
	result = db.Delete(&category)
	return result.Error
}
