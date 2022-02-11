package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	userName := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + userName + " dbname=" + dbName + "sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate()

}
