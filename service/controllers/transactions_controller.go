package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/truecoder34/l0-wb-nats-service/service/models"
	"github.com/truecoder34/l0-wb-nats-service/service/responses"
)

func (server *Server) ParseTransactionsMessage(messageData []byte) (models.Transaction, error) {
	var tr models.Transaction
	err := json.Unmarshal(messageData, &tr)
	if err != nil {
		// TODO : ADD CHECK TO PREVENT UNEXPECTED DATA PROCESSING
		log.Print(err)
	}

	transaction, err := tr.CreatedNestedTransaction(server.DB, tr)
	log.Println(transaction)

	return tr, err
}

func (server *Server) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	transactions, err := transaction.FindAllTransactions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactions)
}

/*
	GET - Get TRANSACTION by its id
		[INPUT] - param ID
*/
func (server *Server) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := uuid.FromString(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	transaction := models.Transaction{}
	transactionReceived, err := transaction.FindTransactionByID(server.DB, tid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactionReceived)
}
