package usecases

import (
	"gomart-api/master/models"
	"gomart-api/master/repositories"
)

type categoryUsecaseImpl struct {
	categoryRepo repositories.CategoryRepository
}

func (c categoryUsecaseImpl) GetAllCategory() ([]*models.Category, error) {
	categorys, err := c.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (c categoryUsecaseImpl) GetCategoryById(id string) (*models.Category, error) {
	category, err := c.categoryRepo.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c categoryUsecaseImpl) AddCategory(categoryName string) error {
	err := c.categoryRepo.AddCategory(categoryName)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryUsecaseImpl) UpdateCategory(categoryID, categoryName string) error {
	err := c.categoryRepo.UpdateCategory(categoryID, categoryName)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryUsecaseImpl) DeleteCategory(id string) error {
	err := c.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

func InitCategoryUsecase(categoryRepo repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecaseImpl{categoryRepo}
}
