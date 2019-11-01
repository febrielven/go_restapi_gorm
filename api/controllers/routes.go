package controllers

import (
	"github.com/febrielven/go_restapi_gorm/api/middlewares"
)

func (s *Server) initiallizeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.GetDetail)).Methods("GET")
}
