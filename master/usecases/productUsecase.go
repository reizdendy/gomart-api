package usecases

import "gomart-api/master/models"

type ProductUsecase interface {
	GetAllProduct() ([]*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	AddProduct(productCode, productName, productCategory string, productPrice int) error
	UpdateProduct(productId, productCode, productName, productCategory string, productPrice int) error
	DeleteProduct(id string) error
}
