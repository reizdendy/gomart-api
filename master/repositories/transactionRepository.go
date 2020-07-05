package repositories

import "gomart-api/master/models"

type TransactionRepository interface {
	GetAllTransaction() ([]*models.Transaction, error)
	GetTransactionById(transactionID string) (*models.Transaction, error)
	AddTransaction(productId string, qty int) error
	UpdateTransaction(transactionID, productId string, qty int) error
	DeleteTransaction(transactionID string) error
}
