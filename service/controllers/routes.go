package controllers

import "github.com/truecoder34/l0-wb-nats-service/service/middlewares"

func (s *Server) initializeRoutes() {
	//base login
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(s.GetTransactions)).Methods("GET")
	s.Router.HandleFunc("/transaction/{id}", middlewares.SetMiddlewareJSON(s.GetTransaction)).Methods("GET")

	s.Router.HandleFunc("/transactions-cache", middlewares.SetMiddlewareJSON(s.GetTransactionsFromCache)).Methods("GET")
	s.Router.HandleFunc("/transaction-cache/{id}", middlewares.SetMiddlewareJSON(s.GetTransactionFromCacheByID)).Methods("GET")

}
