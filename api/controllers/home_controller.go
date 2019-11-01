package controllers

import (
	"fmt"
	"net/http"

	"github.com/febrielven/go_restapi_gorm/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome To This Awesome API")
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
