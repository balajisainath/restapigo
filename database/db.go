package database

import (
	"log"

	"github.com/balajisainath/restapigo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

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
	db.AutoMigrate(models.Book{})
	DB = db

}
