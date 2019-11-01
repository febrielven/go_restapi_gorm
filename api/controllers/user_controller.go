package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/febrielven/go_restapi_gorm/api/models"
	"github.com/febrielven/go_restapi_gorm/api/repositories"
	"github.com/febrielven/go_restapi_gorm/api/responses"
	"github.com/febrielven/go_restapi_gorm/api/utils/formaterror"
	"github.com/gorilla/mux"
)

// // NewPostHandler ...
// func NewPostHandler(server *Server) *userService {

// 	return &userService{
// 		repo: repositories.NewUserRepository(server.DB),
// 	}
// }

// // // // Post ...
// type userService struct {
// 	repo repositories.UserRepository
// }

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userRepository := repositories.NewUserRepository()
	userRepository.Prepare(&user)
	err = userRepository.Validate(&user, "")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreate, err := userRepository.SaveUser(&user)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreate.ID))

	responses.JSON(w, http.StatusCreated, userCreate)

}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	// user := models.User{}
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.FindAllUser()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func (server *Server) GetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	userRepository := repositories.NewUserRepository()

	userGotten, err := userRepository.FindUserByID(&user, uint32(uid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, userGotten)

}

// func UpdateUser(w http.ResponseWriter, r *http.Request)  {
// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseInt(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)

// 	if err !=nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)

// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// }
