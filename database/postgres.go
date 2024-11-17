package database

import (
	"fmt"
	"log"

	"reciepts/config/getconf/getdbconf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dbconf := getdbconf.GetDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbconf.Host, dbconf.Username, dbconf.Password, dbconf.DBname, dbconf.Port)

	var error error
	if db, error = gorm.Open(postgres.Open(dsn), &gorm.Config{}); error != nil {
		log.Fatalf("Error with connecting to database: %s", error.Error())
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Error with connecting to database: %s", err.Error())
	}

	log.Println("Succesfully connected to database")
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if sqlDB, err := db.DB(); err != nil {
		log.Fatalf("Error with closing database: %s", err.Error())
	} else {
		sqlDB.Close()
	}
}
