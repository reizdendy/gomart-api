package usecases

import (
	"gomart-api/master/models"
	"gomart-api/master/repositories"
)

type productUsecaseImpl struct {
	productRepo repositories.ProductRepository
}

func (p productUsecaseImpl) GetAllProduct() ([]*models.Product, error) {
	Products, err := p.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (p productUsecaseImpl) GetProductById(id string) (*models.Product, error) {
	Product, err := p.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return Product, nil
}

func (p productUsecaseImpl) AddProduct(productCode, productName, productCategory string, productPrice int) error {
	err := p.productRepo.AddProduct(productCode, productName, productCategory, productPrice)
	if err != nil {
		return err
	}
	return nil
}

func (p productUsecaseImpl) UpdateProduct(productId, productCode, productName, productCategory string, productPrice int) error {
	err := p.productRepo.UpdateProduct(productId, productCode, productName, productCategory, productPrice)
	if err != nil {
		return err
	}
	return nil
}

func (p productUsecaseImpl) DeleteProduct(id string) error {
	err := p.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func InitProductUsecase(productRepo repositories.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{productRepo}
}
