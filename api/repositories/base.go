package repositories

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Base defines a base respository
type Base struct {
	DB     *gorm.DB
	Router *mux.Router
}
