package repositories

import (
	"database/sql"
	"fmt"
	"gomart-api/master/models"
	"gomart-api/utils"
	"log"
	"time"

	guuid "github.com/google/uuid"
)

type TransactionRepoImpl struct {
	db *sql.DB
}

func (t *TransactionRepoImpl) GetAllTransaction() ([]*models.Transaction, error) {
	rows, err := t.db.Query(utils.GET_ALL_TRANSACTION)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Transactions []*models.Transaction
	for rows.Next() {
		Transaction := models.Transaction{}
		err := rows.Scan(&Transaction.TransactionId, &Transaction.TransactionDate, &Transaction.TransactionProductId, &Transaction.Quantity)

		if err != nil {
			return nil, err
		}

		Transactions = append(Transactions, &Transaction)
	}
	return Transactions, nil
}

func (t *TransactionRepoImpl) GetTransactionById(id string) (*models.Transaction, error) {
	Transaction := models.Transaction{}
	fmt.Println(id)
	err := t.db.QueryRow(utils.GET_TRANSACTION_BY_ID, id).Scan(&Transaction.TransactionId, &Transaction.TransactionDate, &Transaction.TransactionProductId, &Transaction.Quantity)
	if err != nil {
		return nil, err
	}
	return &Transaction, nil
}

func (t *TransactionRepoImpl) AddTransaction(productId string, qty int) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := t.db.Prepare(utils.ADD_TRANSACTION)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(guuid.New(), time.Now().Format("2006-02-01 15:04:05"), productId, qty)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// UpdateTransaction is a function to update Transaction data
func (t *TransactionRepoImpl) UpdateTransaction(transactionId, productId string, qty int) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := t.db.Prepare(utils.UPDATE_TRANSACTION)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(productId, qty, transactionId); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// DeleteTransaction is a function to delete Transaction data
func (t *TransactionRepoImpl) DeleteTransaction(id string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := t.db.Prepare(utils.DELETE_TRANSACTION)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}

	// _, err = data.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return tx.Commit()
}

func InitTransactionRepositoryImpl(db *sql.DB) TransactionRepository {
	return &TransactionRepoImpl{db}
}
