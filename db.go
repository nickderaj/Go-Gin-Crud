package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database.")
	}
}

func MigrateDB() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate Database.")
	}
}
