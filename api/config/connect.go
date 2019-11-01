package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var (
	db *gorm.DB
)

func Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

		db, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			fmt.Printf("Connot connect to %s database", Dbdriver)
			log.Fatal("this is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
