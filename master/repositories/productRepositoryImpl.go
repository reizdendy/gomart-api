package repositories

import (
	"database/sql"
	"gomart-api/master/models"
	"gomart-api/utils"
	"log"

	guuid "github.com/google/uuid"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (p *ProductRepoImpl) GetAllProduct() ([]*models.Product, error) {
	rows, err := p.db.Query(utils.GET_ALL_PRODUCT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Products []*models.Product
	for rows.Next() {
		Product := models.Product{}
		err := rows.Scan(&Product.ProductID, &Product.ProductCode, &Product.ProductName, &Product.ProductCategory, &Product.ProductPrice, &Product.Status)

		if err != nil {
			return nil, err
		}

		Products = append(Products, &Product)
	}
	return Products, nil
}

func (p *ProductRepoImpl) GetProductById(id string) (*models.Product, error) {
	Product := models.Product{}
	err := p.db.QueryRow(utils.GET_PRODUCT_BY_ID, id).Scan(&Product.ProductID, &Product.ProductCode, &Product.ProductName, &Product.ProductCategory, &Product.ProductPrice, &Product.Status)
	if err != nil {
		return nil, err
	}
	return &Product, nil
}

func (p *ProductRepoImpl) AddProduct(productCode, productName, productCategory string, productPrice int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := p.db.Prepare(utils.ADD_PRODUCT)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	data, err := stmt.Exec(guuid.New(), productCode, productName, productCategory, productPrice)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	_, err = data.LastInsertId()
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// UpdateProduct is a function to update Product data
func (p *ProductRepoImpl) UpdateProduct(productId, productCode, productName, productCategory string, productPrice int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := p.db.Prepare(utils.UPDATE_PRODUCT)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(productCode, productName, productCategory, productPrice, productId); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// DeleteProduct is a function to delete Product data
func (p *ProductRepoImpl) DeleteProduct(id string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := p.db.Prepare(utils.DELETE_PRODUCT)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	data, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}

	_, err = data.LastInsertId()
	if err != nil {
		return err
	}
	return tx.Commit()
}

func InitProductRepositoryImpl(db *sql.DB) ProductRepository {
	return &ProductRepoImpl{db}
}
