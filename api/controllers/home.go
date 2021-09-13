package controllers

import (
	"net/http"
	
	"blog/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to the api")
}