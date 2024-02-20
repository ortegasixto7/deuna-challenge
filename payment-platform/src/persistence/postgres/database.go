package postgres

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var totalRetries = 0

func Init() {
	time.Sleep(2 * time.Second)
	databaseUrl := os.Getenv("DATABASE_URL")
	var GO_ENV string = os.Getenv("GO_ENV")
	if GO_ENV == "docker" {
		databaseUrl = os.Getenv("DATABASE_URL_DOCKER")
	}
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting Database, retrying again")
		fmt.Println(err)
		if totalRetries < 3 {
			totalRetries++
			Init()
		}
		panic("Error In Database Connection")
	}

	Database = db

}
