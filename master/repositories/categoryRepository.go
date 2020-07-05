package repositories

import "gomart-api/master/models"

type CategoryRepository interface {
	GetAllCategory() ([]*models.Category, error)
	GetCategoryById(categoryID string) (*models.Category, error)
	AddCategory(categoryName string) error
	UpdateCategory(categoryID, categoryName string) error
	DeleteCategory(categoryID string) error
}
