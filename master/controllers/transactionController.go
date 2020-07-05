package controllers

import (
	"encoding/json"
	"fmt"
	"gomart-api/master/models"
	"gomart-api/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type transactionHandler struct {
	transactionUsecase usecases.TransactionUsecase
}

var (
	transaction         models.Transaction
	transactionResponse models.Response
)

func GetMuxtransaction(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func TransactionController(r *mux.Router, service usecases.TransactionUsecase) {
	transactionHandler := transactionHandler{service}
	r.HandleFunc("/transactions", transactionHandler.ListTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id}", transactionHandler.ListTransactionByID).Methods(http.MethodGet)
	r.HandleFunc("/transaction", transactionHandler.AddTransaction).Methods(http.MethodPost)
	r.HandleFunc("/transaction", transactionHandler.UpdateTransaction).Methods(http.MethodPut)
	r.HandleFunc("/transaction/{id}", transactionHandler.DeleteTransaction).Methods(http.MethodDelete)
}

func (p *transactionHandler) ListTransaction(w http.ResponseWriter, r *http.Request) {
	transactions, err := p.transactionUsecase.GetAllTransaction()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		transactionResponse.Messages = "List Of Transaction Data"
		transactionResponse.Data = transactions
		transactionResponse.Status = http.StatusOK
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.Write([]byte("Data Not Found"))
		}
		byteOftransaction, err := json.Marshal(transactionResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOftransaction)
	}
}

func (p *transactionHandler) ListTransactionByID(w http.ResponseWriter, r *http.Request) {
	idtransaction := GetMuxtransaction("id", r)
	fmt.Println(idtransaction)
	transaction, err := p.transactionUsecase.GetTransactionById(idtransaction)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		transactionResponse.Messages = "List Of Transaction Data"
		transactionResponse.Data = transaction
		transactionResponse.Status = http.StatusOK
		byteOftransaction, err := json.Marshal(transactionResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOftransaction)
	}
}

//AddTransaction is a controller of AddAtransaction service (Insert / POST). It got the json data from Request Body RAW
func (p *transactionHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	err := p.transactionUsecase.AddTransaction(transaction.TransactionProductId, transaction.Quantity)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		transactionResponse.Messages = "Insert transaction Data Success!"
		transactionResponse.Status = http.StatusOK
		transactionResponse.Data = ""
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		byteOftransaction, err := json.Marshal(transactionResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOftransaction)
	}
}

//UpdateTransaction is a controller of UpdateAtransaction service (Update/PUT). It got the json data from Request Body RAW
func (p *transactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	err := p.transactionUsecase.UpdateTransaction(transaction.TransactionId, transaction.TransactionProductId, transaction.Quantity)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {

		transactionResponse.Messages = "Update transaction Success"
		transactionResponse.Status = http.StatusOK
		transactionResponse.Data = ""
		byteOftransaction, err := json.Marshal(transactionResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOftransaction)
	}
}

// DeleteTransaction is a controller of  DeleteAtransaction service  (DeleTE). It got the json data from Params
func (s *transactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idtransaction := GetMuxtransaction("id", r)

	err := s.transactionUsecase.DeleteTransaction(idtransaction)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		transactionResponse.Messages = "Delete transaction Success"
		transactionResponse.Status = http.StatusOK
		transactionResponse.Data = ""
		byteOftransaction, err := json.Marshal(transactionResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOftransaction)
	}
}
