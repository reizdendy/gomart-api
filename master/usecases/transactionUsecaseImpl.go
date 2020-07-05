package usecases

import (
	"gomart-api/master/models"
	"gomart-api/master/repositories"
)

type transactionUsecaseImpl struct {
	TransactionRepo repositories.TransactionRepository
}

func (p transactionUsecaseImpl) GetAllTransaction() ([]*models.Transaction, error) {
	transactions, err := p.TransactionRepo.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p transactionUsecaseImpl) GetTransactionById(id string) (*models.Transaction, error) {
	Transaction, err := p.TransactionRepo.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	return Transaction, nil
}

func (p transactionUsecaseImpl) AddTransaction(productId string, qty int) error {
	err := p.TransactionRepo.AddTransaction(productId, qty)
	if err != nil {
		return err
	}
	return nil
}

func (p transactionUsecaseImpl) UpdateTransaction(transactionId, productId string, qty int) error {
	err := p.TransactionRepo.UpdateTransaction(transactionId, productId, qty)
	if err != nil {
		return err
	}
	return nil
}

func (p transactionUsecaseImpl) DeleteTransaction(id string) error {
	err := p.TransactionRepo.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func InitTransactionUsecase(TransactionRepo repositories.TransactionRepository) TransactionUsecase {
	return &transactionUsecaseImpl{TransactionRepo}
}
