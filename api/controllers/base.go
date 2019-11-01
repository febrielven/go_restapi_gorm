package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febrielven/go_restapi_gorm/api/config"
	"github.com/febrielven/go_restapi_gorm/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// func NewPostHandler(server *Server) *Server {

// 	return &Server{
// 		repo: repositories.NewUserRepository(server.base.DB),
// 	}
// }

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	config.Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)
	server.DB = config.GetDB()
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initiallizeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 5050")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
