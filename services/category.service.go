package services

import (
	"errors"

	"kasir-api/models"
	"kasir-api/repositories"
)

// CategoryService handles business logic for categories
type CategoryService struct {
	repo repositories.ICategoryRepository
}

// NewCategoryService creates a new category service instance
func NewCategoryService(repo repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// CreateCategory creates a new category in the database
func (s *CategoryService) CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	return s.repo.Create(category)
}

// GetAllCategories retrieves all categories from the database
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

// GetCategoryByID retrieves a single category by ID
func (s *CategoryService) GetCategoryByID(id string) (*models.Category, error) {
	return s.repo.FindByID(id)
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(id string, updateData *models.Category) (*models.Category, error) {
	// Validate required fields
	if updateData.Name == "" {
		// Get current category to check if it has a name
		current, err := s.repo.FindByID(id)
		if err != nil {
			return nil, err
		}
		if current.Name == "" {
			return nil, errors.New("category name is required")
		}
	}

	return s.repo.Update(id, updateData)
}

// DeleteCategory deletes a category by ID (soft delete)
func (s *CategoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
