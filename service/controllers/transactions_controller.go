package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/truecoder34/l0-wb-nats-service/service/models"
	"github.com/truecoder34/l0-wb-nats-service/service/responses"
)

func (server *Server) CreateTransactionFromNATS(messageData []byte) (models.Transaction, error) {
	var tr models.Transaction
	err := json.Unmarshal(messageData, &tr)
	if err != nil {
		// TODO : ADD CHECK TO PREVENT UNEXPECTED DATA PROCESSING
		log.Print(err)
	}

	// create to DB
	transaction, err := tr.CreatedNestedTransaction(server.DB, tr)
	log.Println(transaction)

	// add to cache
	server.cache.Set(transaction.ID.String(), *transaction, 50000*time.Minute)
	//trCahce, i := server.cache.Get("testKey")
	//log.Println("return from cache : ",trCahce, i)

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

/*
	GET TRANSACTIONS FROM CACHE
*/
func (server *Server) GetTransactionsFromCache(w http.ResponseWriter, r *http.Request) {
	transactions, err := server.cache.GetAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactions)
}

/*
	GET TRANSACTION FROM CACHE BY ID
*/
func (server *Server) GetTransactionFromCacheByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := uuid.FromString(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	transactionReceived, found := server.cache.Get(tid.String())
	if found == false {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactionReceived)
}
