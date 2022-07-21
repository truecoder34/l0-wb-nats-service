package controllers

import "github.com/truecoder34/l0-wb-nats-service/service/middlewares"

func (s *Server) initializeRoutes() {
	//base login
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

}
