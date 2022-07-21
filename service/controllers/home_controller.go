package controllers

import (
	"net/http"

	"github.com/truecoder34/l0-wb-nats-service/service/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To API")

}
